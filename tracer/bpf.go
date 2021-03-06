// vim: set syntax=c:
package tracer

const src = `
#include <linux/sched.h>
#include <uapi/linux/ptrace.h>

BPF_PERF_OUTPUT(events);

enum event_type {
	EVENT_START = 1,
	EVENT_FINISH,
};

struct data_t {
	enum event_type type;

	char comm[TASK_COMM_LEN];
	int  exitcode;
	u32  pid, ppid;
	u64  ts;
};

static inline void
__submit(struct pt_regs* ctx, struct data_t* data)
{
	struct task_struct* task;

	task = (struct task_struct*)bpf_get_current_task();

	data->ts   = bpf_ktime_get_ns();
	data->pid  = task->tgid;
	data->ppid = task->real_parent->tgid;
	bpf_get_current_comm(&data->comm, sizeof(data->comm));

	events.perf_submit(ctx, data, sizeof(*data));
}

static inline void
__submit_start(struct pt_regs* ctx)
{
	struct data_t data = {};

	data.type = EVENT_START;

	__submit(ctx, &data);
}

static inline void
__submit_finish(struct pt_regs* ctx, int code)
{
	struct data_t data = {};

	data.type     = EVENT_FINISH;
	data.exitcode = code;

	__submit(ctx, &data);
}

int
kr__sys_execve(struct pt_regs* ctx,
               const char __user* filename,
               const char __user* const __user* __argv,
               const char __user* const __user* __envp)
{
	int ret = PT_REGS_RC(ctx);
	if (ret != 0) {
		return 0;
	}

	__submit_start(ctx);

	return 0;
}

int
k__do_exit(struct pt_regs* ctx, long code)
{
	__submit_finish(ctx, code);
	return 0;
}
`
