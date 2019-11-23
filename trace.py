#!/usr/bin/python

# trace.py  Traces the execution of processes.
#
# [execve]: https://elixir.bootlin.com/linux/v5.0/source/fs/exec.c#L1963
# [close]: 

from __future__ import print_function
from bcc import BPF
from pprint import pprint


program = BPF(text=r"""
#include <linux/sched.h>
#include <uapi/linux/ptrace.h>

#define ARG_LEN 128

BPF_PERF_OUTPUT(events);

enum event_type
{
        EVENT_START = 1,
        EVENT_FINISH,
};

struct data_t
{
        u32 pid, ppid;
        u64 ts;
        char comm[TASK_COMM_LEN];
        char argv[ARG_LEN];
        enum event_type type;
};


static inline void
__submit(struct pt_regs* ctx, enum event_type type)
{
        struct task_struct* task;
        struct data_t data = {};

        task = (struct task_struct*)bpf_get_current_task();

        data.type = type;
        data.ts = bpf_ktime_get_ns();
        data.pid = task->tgid;
        data.ppid = task->real_parent->tgid;
        bpf_get_current_comm(&data.comm, sizeof(data.comm));

        events.perf_submit(ctx, &data, sizeof(data));
}

int
syscall__execve(struct pt_regs* ctx,
                const char __user* filename,
                const char __user* const __user* __argv,
                const char __user* const __user* __envp)
{
        bpf_trace_printk("start\n");

        __submit(ctx, EVENT_START);

        return 0;
}

int
kprobe__do_exit(struct pt_regs* ctx, long code)
{
        bpf_trace_printk("finish\n");

        __submit(ctx, EVENT_FINISH);

        return 0;
}
""")

program.attach_kprobe(
        event=program.get_syscall_fnname("execve"),
        fn_name="syscall__execve")

program.attach_kprobe(
        event="do_exit",
        fn_name="kprobe__do_exit")


class EventType(object):
    EVENT_START = 1
    EVENT_FINISH = 2


procs = {}

def handle_events(cpu, data, size):
    event = program["events"].event(data)

    if event.type == EventType.EVENT_START:
        procs[event.pid] = event
        print("start")

    elif event.type == EventType.EVENT_FINISH:
        if not event.pid in procs:
            return

        seconds_elapsed = (event.ts - procs[event.pid].ts) / (10**9)

        print("finished %s in %3fs" % (event.comm, seconds_elapsed))
        del procs[event.pid]
    


program["events"].open_perf_buffer(handle_events)

while 1:
    try:
        program.perf_buffer_poll()
    except KeyboardInterrupt:
        exit()
