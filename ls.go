package main

import "flag"

type (
	input struct {
		Flag_1    bool
		Flag_a    bool
		Flag_A    bool
		Flag_c    bool
		Flag_d    bool
		Flag_f    bool
		Flag_F    bool
		Flag_h    bool
		Flag_H    bool
		Flag_i    bool
		Flag_l    bool
		Flag_L    bool
		Flag_n    bool
		Flag_p    bool
		Flag_q    bool
		Flag_r    bool
		Flag_R    bool
		Flag_u    bool
		Flag_U    bool
		Flag_sort string
	}
)

func getInputData() *input {

	var (
		aflag bool
		cflag bool
		fflag bool
		lflag bool
		nflag bool
		uflag bool
		Uflag bool
		sort  string
	)

	singleColumnFlagPtr := flag.Bool("1", false, "force output to be one per line")
	allButHereAndAboveFlagPtr := flag.Bool("A", false, "List all entries except for '.' and '..'.")
	allFlagPtr := flag.Bool("a", false, "Show hidden files (those beginning with '.').")
	creationTimeFlagPtr := flag.Bool("c", false, "Use time file's status was last changed instead of last modification time for sorting or printing.")
	dirsAsFilesFlagPtr := flag.Bool("d", false, "If argument is ")
	noSortFlagPtr := flag.Bool("f", false, "Like -U but turns on -a and disables -r, -S and -t.")
	typeFlagPtr := flag.Bool("F", false, "Append a file type indicator to all special files.")
	symlinkFlagPrt := flag.Bool("H", false, "List information about the targets of symbolic links specified on the command line instead of the links themselves.")
	humanFlagPtr := flag.Bool("h", false, "Show filesizes in human-readable format.")
	indexFlagPtr := flag.Bool("i", false, "Print the index number of each file.")
	targetsFlagPtr := flag.Bool("L", false, "List information about the targets of symbolic links instead of the links themselves.")
	longFlagPtr := flag.Bool("l", false, "List detailed information about each file, including their type, permissions, links, owner, group, size or major and minor numbers if the file is a character/block device, and last file status/modification time.")
	longNumericFlagPtr := flag.Bool("n", false, "List detailed information about each file, including their type, permissions, links, owner, group, size or major and minor numbers if the file is a character/block device, and last file status/modification time, but with numeric IDs.")
	dirTypeFlagPtr := flag.Bool("p", false, "Append a file type indicator to directories.")
	printNonPrintablesPtr := flag.Bool("q", false, "Replace non-printable characters in filenames with '?'.")
	recurseFlagPtr := flag.Bool("R", false, "List directory content recursively.  The -1 flag is set implicitly.")
	reverseSortFlagPtr := flag.Bool("r", false, "Reverse the sort order.")
	sortBySizePtr := flag.Bool("S", false, "Sort files by size (in decreasing order).")
	sortByTimePtr := flag.Bool("t", false, "Sort files by last file status/modification time instead of by name.")
	unsortedPtr := flag.Bool("U", false, "Keep the list unsorted.")
	accessTimePtr := flag.Bool("u", false, "Use file's last access time instead of last modification time for sorting or printing.")

	flag.Parse()

	if *allFlagPtr {
		aflag = true
	}

	if *creationTimeFlagPtr {
		cflag = true
		uflag = false
	}

	if *noSortFlagPtr {
		aflag = true
		fflag = true
		Uflag = true
	}

	if *longFlagPtr {
		lflag = true
	}

	if *longNumericFlagPtr {
		lflag = true
		nflag = true
	}

	if *accessTimePtr {
		uflag = true
		cflag = false
	}

	if *unsortedPtr {
		Uflag = true
	}

	if *sortBySizePtr {
		sort = "S"
	} else if *sortByTimePtr {
		sort = "t"
	} else {
		sort = "n"
	}

	return &input{
		Flag_1:    *singleColumnFlagPtr,
		Flag_a:    aflag,
		Flag_A:    *allButHereAndAboveFlagPtr,
		Flag_c:    cflag,
		Flag_d:    *dirsAsFilesFlagPtr,
		Flag_f:    fflag,
		Flag_F:    *typeFlagPtr,
		Flag_h:    *humanFlagPtr,
		Flag_H:    *symlinkFlagPrt,
		Flag_i:    *indexFlagPtr,
		Flag_l:    lflag,
		Flag_L:    *targetsFlagPtr,
		Flag_n:    nflag,
		Flag_p:    *dirTypeFlagPtr,
		Flag_q:    *printNonPrintablesPtr,
		Flag_r:    *reverseSortFlagPtr,
		Flag_R:    *recurseFlagPtr,
		Flag_u:    uflag,
		Flag_U:    Uflag,
		Flag_sort: sort,
	}
}
