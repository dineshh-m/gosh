package term

const prefix = "➜  "
const (
  green = "\x1b[32m"
  cyan = "\x1b[36m"
  defaultColor = "\x1b[0m"
)

func Term(cwd string) string {
  return GreenText(prefix) + CyanText(cwd + " ")
}

func CyanText(text string) string {
  return cyan + text + defaultColor
}

func GreenText(text string) string {
  return green + text + defaultColor
}
