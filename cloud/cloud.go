package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb { //создание облачной бд(файла)
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) { // Методы - читаем файл указанный
	return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) { // метод для записи данных в файл - принимает []byte

}
