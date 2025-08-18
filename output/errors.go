package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	intVal, ok := value.(int)
	if ok {
		color.Red("Код ошибки %d", intVal)
		return
	}
	strVal, ok := value.(string)

	if ok {
		color.Red(strVal)
		return
	}
	errorVal, ok := value.(error)
	if ok {
		color.Red(errorVal.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")

}
