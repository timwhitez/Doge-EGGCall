# Doge-EGGCall
Like Direct Syscall but more EGG :)

supports syscalls instruction replacement with an EGG (to be dynamically replaced)

```
eggcall.MemHgate()
eggreplace.FindAndReplace(
		[]byte{0x65,0x67,0x67,0x63,0x61,0x6c,0x6c},
		[]byte{0x90,0x90,0x0f,0x05,0x90,0x90,0x90},
		reflect.ValueOf(eggcall.HgSyscall).Pointer())
eggcall.HgSyscall()

```


### ref
https://klezvirus.github.io/RedTeaming/AV_Evasion/NoSysWhisper/
