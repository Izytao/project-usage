package tools

// RedirectStderr 在 Unix/Linux 系统上，你可以将 stderr 重定向到一个文件，以便将错误消息记录到文件中而不是终端上。
//func RedirectStderr(f *os.File) {
//	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
//	if err != nil {
//		log.Printf("Failed to redirect stderr to file: %v", err)
//	}
//}
