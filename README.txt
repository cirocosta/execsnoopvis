

TYPES

	type Node struct {
		Pid uint

		Argv     []string
		Command  string
		ExitCode int
		Ppid     uint

		Parent   *Node
		Children []*Node
	}



SCANNING

	// given a reader, parse the lines and create the list of per-pid
	// mapping of those entries


	var perPidNodeMapping = map[uint]*Node

	for scanner.Scan() {
		line := scanner.Text()

		node, err := ParseNode(line)
		if err != nil {
			return
		}

		perPidNodeMapping[node.Pid] = node
	}

	// fill the `*Parent` pointers
	// fill the `*Parent` pointers
	for _, node := range perPidNodeMapping {
		node.Parent = perPidNodeMapping[node.Ppid]

		if node.Parent != nil {
			node.Parent.Children = append(
				node.Parent.Children, 
				node,
			)
		}
	}




ROOTS DISCOVERY

	// for every node that we have discovered, find those whose `parent` are
	// `nil`

	for _, node := range 




EXAMPLE INPUT


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
	jump             17276  16232    0 /usr/bin/jump chdir
	sh               17284  17283    0 /bin/sh -c command -v debian-sa1 > /dev/nu...
	debian-sa1       17285  17284    0 /usr/lib/sysstat/debian-sa1 1 1
	jump             17286  16232    0 /usr/bin/jump chdir
	ls               17291  16232    0 /bin/ls --color=auto
	jump             17292  16232    0 /usr/bin/jump chdir


DOT NOTATION

	
	digraph G {

		subgraph G0 {
			color=gray
		}


		subgraph G0 {
			color=gray
		}
	}

