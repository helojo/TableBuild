package libs

type ConsoleLevel int

// iota 初始化后会自动递增
const (
    Log ConsoleLevel = iota
    Info
    Debug
    Warn
    Error
)

func Colorize(text string, status ConsoleLevel) string {
    out := ""
    switch status {
        case Info:
            out = "\033[32;1m"    // Blue
        case Error:
            out = "\033[31;1m"    // Red
        case Warn:
            out = "\033[33;1m"    // Yellow
        case Debug:
            out = "\033[34;1m"    // Green
        default:
            out = "\033[0m"    // Default
    }
    return out+text+"\033[0m"
}