package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type logger struct {
	*zap.Logger
}

// シングルトン
var singleton = newLogger()

func newLogger() *logger {
	zap, _ := zap.NewDevelopment()
	return &logger{zap}
}

// Debugログ
func Debug(msg string) {
	singleton.Debug(msg)
}

// Infoログ
func Info(msg string) {
	singleton.Info(msg)
}

// Warnログ
func Warn(msg string) {
	singleton.Warn(msg)
}

// Errorログ
func Error(msg string) {
	singleton.Error(msg)
}

// TODO 開発時のみ表示する
func StartLog() {
	fmt.Println("          _____                    _____                    _____                    _____")
	fmt.Println("         /\\    \\                  /\\    \\                  /\\    \\                  /\\    \\")
	fmt.Println("        /::\\    \\                /::\\    \\                /::\\    \\                /::\\    \\")
	fmt.Println("       /::::\\    \\              /::::\\    \\              /::::\\    \\              /::::\\    \\")
	fmt.Println("      /::::::\\    \\            /::::::\\    \\            /::::::\\    \\            /::::::\\    \\")
	fmt.Println("     /:::/\\:::\\    \\          /:::/\\:::\\    \\          /:::/\\:::\\    \\          /:::/\\:::\\    \\")
	fmt.Println("    /:::/  \\:::\\    \\        /:::/__\\:::\\    \\        /:::/__\\:::\\    \\        /:::/  \\:::\\    \\")
	fmt.Println("   /:::/    \\:::\\    \\      /::::\\   \\:::\\    \\      /::::\\   \\:::\\    \\      /:::/    \\:::\\    \\")
	fmt.Println("  /:::/    / \\:::\\    \\    /::::::\\   \\:::\\    \\    /::::::\\   \\:::\\    \\    /:::/    / \\:::\\    \\")
	fmt.Println(" /:::/    /   \\:::\\ ___\\  /:::/\\:::\\   \\:::\\____\\  /:::/\\:::\\   \\:::\\____\\  /:::/    /   \\:::\\    \\")
	fmt.Println("/:::/____/  ___\\:::|    |/:::/  \\:::\\   \\:::|    |/:::/  \\:::\\   \\:::|    |/:::/____/     \\:::\\____\\")
	fmt.Println("\\:::\\    \\ /\\  /:::|____|\\::/   |::::\\  /:::|____|\\::/    \\:::\\  /:::|____|\\:::\\    \\      \\::/    /")
	fmt.Println(" \\:::\\    /::\\ \\::/    /  \\/____|:::::\\/:::/    /  \\/_____/\\:::\\/:::/    /  \\:::\\    \\      \\/____/")
	fmt.Println("  \\:::\\   \\:::\\ \\/____/         |:::::::::/    /            \\::::::/    /    \\:::\\    \\")
	fmt.Println("   \\:::\\   \\:::\\____\\           |::|\\::::/    /              \\::::/    /      \\:::\\    \\")
	fmt.Println("    \\:::\\  /:::/    /           |::| \\::/____/                \\::/____/        \\:::\\    \\")
	fmt.Println("     \\:::\\/:::/    /            |::|   |                                        \\:::\\    \\")
	fmt.Println("      \\::::::/    /             |::|   |                                         \\:::\\    \\")
	fmt.Println("       \\::::/    /              \\::|   |                                          \\:::\\____\\")
	fmt.Println("        \\::/____/                \\:|   |                                           \\::/    /")
	fmt.Println("                                  \\|___|                                            \\/____/")
	fmt.Println("")
}
