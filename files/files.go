package files

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type Jsondb struct { //Структура jsondb - нужна для того чтобы указывать
	filename string // в какую бд(джсон файл в этом примере) мы хотим записать файл
}

func NewJsonDb(name string) *Jsondb { //Создание файла с именем
	return &Jsondb{ //Возвращение указателя на jsondb c изменённым именем
		filename: name, //ввод в параметр filename получаемого значения name
	}
}

func (db *Jsondb) Read() ([]byte, error) { // Методы - читаем файл указанный
	data, err := os.ReadFile(db.filename) //читаем файл с именем filename
	if err != nil {
		return nil, err
	}
	return data, nil //возвращаем информацию полученную при чтении этого файла - тип []byte
}

func (db *Jsondb) Write(content []byte) { // метод для записи данных в файл - принимает []byte
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	color.Green("Успешная запись")
}
