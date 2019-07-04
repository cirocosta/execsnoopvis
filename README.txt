



TYPES

	type Row struct {
		Command  string
		Pid      uint
		Ppid     uint
		ExitCode int
		Argv     []string
	}

	type Node struct {
		Parent *Node
		Row
	}



SCANNING

	var (
		perPidRowMapping = map[uint]*Row
	)

	for scanner.Scan() {

		line := scanner.Text()

		row, err := ParseRow(line)
		if err != nil {
			return
		}

		perPidRowMapping[row.Pid] = row
	}



FOREST FORMATION

	for pid, row := range perPidRowMapping {

	}






PCOMM            PID    PPID   RET ARGS
go               17123  16232    0 /usr/local/go/bin/go build -v .
cgo              17130  17123    0 /usr/local/go/pkg/tool/linux_amd64/cgo -V=full
compile          17131  17123    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
compile          17142  17123    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
compile          17137  17123    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
compile          17147  17123    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
asm              17158  17123    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17160  17123    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17164  17123    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17169  17123    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
link             17178  17123    0 /usr/local/go/pkg/tool/linux_amd64/link -V=full
jump             17184  16232    0 /usr/bin/jump chdir
go               17189  16232    0 /usr/local/go/bin/go clean
jump             17196  16232    0 /usr/bin/jump chdir
go               17201  16232    0 /usr/local/go/bin/go build -v .
cgo              17210  17201    0 /usr/local/go/pkg/tool/linux_amd64/cgo -V=full
compile          17211  17201    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
compile          17213  17201    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
compile          17212  17201    0 /usr/local/go/pkg/tool/linux_amd64/compile -V=full
asm              17231  17201    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17230  17201    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17232  17201    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
asm              17247  17201    0 /usr/local/go/pkg/tool/linux_amd64/asm -V=full
link             17252  17201    0 /usr/local/go/pkg/tool/linux_amd64/link -V=full
link             17257  17201    0 /usr/local/go/pkg/tool/linux_amd64/link -o /tmp/go-build487788134/b001/exe/a.out -importcfg /tmp/go-build487788134/b001/importcfg.link -buildmode=exe -buildid=kIitCVFqEpKt7i0cOSWc/i1wzLezL9n7SWYOAVsl0/8DDDY0koOx9LmVh80Tkd/kIitCVFqEpKt7i0cOSWc -extld=gcc /home/ubuntu/.cache/go-build/a7/a71826a7dbe3d00c2b9d52fa68e3ef110efc31d87fdc2cdd2d21f4ae3eec0813-d
gcc              17262  17257    0 /usr/bin/gcc -Wl,--compress-debug-sections=zlib-gnu trivial.c
cc1              17263  17262    0 /usr/lib/gcc/x86_64-linux-gnu/7/cc1 -quiet -imultiarch x86_64-linux-gnu trivial.c -quiet -dumpbase trivial.c -mtune=generic -march=x86-64 -auxbase trivial -fstack-protector-strong -Wformat -Wformat-security -o /tmp/ccVvmduh.s
as               17264  17262    0 /usr/bin/as --64 -o /tmp/ccL6nLft.o /tmp/ccVvmduh.s
collect2         17265  17262    0 /usr/lib/gcc/x86_64-linux-gnu/7/collect2 -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/ccfa511E.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -dynamic-linker /lib64/ld-linux-x86-64.so.2 -pie ...
ld               17266  17265    0 /usr/bin/ld -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/ccfa511E.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -dynamic-linker /lib64/ld-linux-x86-64.so.2 -pie ...
gcc              17268  17257    0 /usr/bin/gcc -no-pie trivial.c
cc1              17269  17268    0 /usr/lib/gcc/x86_64-linux-gnu/7/cc1 -quiet -imultiarch x86_64-linux-gnu trivial.c -quiet -dumpbase trivial.c -mtune=generic -march=x86-64 -auxbase trivial -fstack-protector-strong -Wformat -Wformat-security -o /tmp/cc74sMak.s
as               17270  17268    0 /usr/bin/as --64 -o /tmp/ccBZgQ0v.o /tmp/cc74sMak.s
collect2         17271  17268    0 /usr/lib/gcc/x86_64-linux-gnu/7/collect2 -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/cc9RGYRH.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -dynamic-linker /lib64/ld-linux-x86-64.so.2 -z ...
ld               17272  17271    0 /usr/bin/ld -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/cc9RGYRH.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -dynamic-linker /lib64/ld-linux-x86-64.so.2 -z ...
gcc              17273  17257    0 /usr/bin/gcc -m64 -o /tmp/go-build487788134/b001/exe/a.out -rdynamic -Wl,--compress-debug-sections=zlib-gnu /tmp/go-link-533938121/go.o /tmp/go-link-533938121/000000.o /tmp/go-link-533938121/000001.o /tmp/go-link-533938121/000002.o /tmp/go-link-533938121/000003.o /tmp/go-link-533938121/000004.o /tmp/go-link-533938121/000005.o /tmp/go-link-533938121/000006.o /tmp/go-link-533938121/000007.o /tmp/go-link-533938121/000008.o /tmp/go-link-533938121/000009.o /tmp/go-link-533938121/000010.o /tmp/go-link-533938121/000011.o /tmp/go-link-533938121/000012.o ...
collect2         17274  17273    0 /usr/lib/gcc/x86_64-linux-gnu/7/collect2 -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/ccIFvyxm.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -export-dynamic -dynamic-linker /lib64/ld-linux-x86-64.so.2 ...
ld               17275  17274    0 /usr/bin/ld -plugin /usr/lib/gcc/x86_64-linux-gnu/7/liblto_plugin.so -plugin-opt=/usr/lib/gcc/x86_64-linux-gnu/7/lto-wrapper -plugin-opt=-fresolution=/tmp/ccIFvyxm.res -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s -plugin-opt=-pass-through=-lc -plugin-opt=-pass-through=-lgcc -plugin-opt=-pass-through=-lgcc_s --sysroot=/ --build-id --eh-frame-hdr -m elf_x86_64 --hash-style=gnu --as-needed -export-dynamic -dynamic-linker /lib64/ld-linux-x86-64.so.2 ...
jump             17276  16232    0 /usr/bin/jump chdir
sh               17284  17283    0 /bin/sh -c command -v debian-sa1 > /dev/null && debian-sa1 1 1
debian-sa1       17285  17284    0 /usr/lib/sysstat/debian-sa1 1 1
jump             17286  16232    0 /usr/bin/jump chdir
ls               17291  16232    0 /bin/ls --color=auto
jump             17292  16232    0 /usr/bin/jump chdir
