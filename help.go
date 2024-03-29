//Gin-Gonic
c.Param() - //возвращает значение параметра URL
c.FullPath()-//возвращает полный путь совпавшего маршрута
c.String()-//записывает заданную строку в тело ответа.
c.DefaultQuery()-//возвращает значение запроса URL-адреса с ключом,
//если оно существует, в противном случае возвращает указанную строку
c.Query()-//возвращает значение запроса URL с ключом,
//если оно существует, в противном случае возвращает пустую строку
c.PostForm()-//возвращает указанный ключ из POST или многостраничной формы
c.DefaultPostForm() -//возвращает указанный ключ из  POST или составной формы,
// если она существует, в противном случае она возвращает указанную строку
c.JSON()-//сериализует данную структуру как JSON в тело ответа
c.QueryMap()-//возвращает карту для заданного ключа запроса.
c.PostFormMap()-//возвращает карту для заданного ключа формы.
c.FormFile()-//возвращает первый файл для предоставленного ключа формы
c.SaveUploadedFile() - //загружает файл в конкретный dst
c.MultipartForm() - //проанализированная форма, загрузку файлов.
c.ShouldBind() - /*&ShouldBind проверяет Content-Type для автоматического выбора механизма привязки. В зависимости от заголовка Content-Type используются разные привязки:
"application / json" -> привязка JSON
"application / xml" -> привязка XML
в противном случае -> возвращает ошибку*/
c.String() - //записывает заданную строку в тело ответа.
c.Bind() - //проверяет Content-Type для автоматического выбора механизма привязки.
// В зависимости от заголовка Content-Type используются разные привязки
//"application/json" --> JSON binding
//"application/xml"  --> XML binding

gin.DisableConsoleColor()// Создание и запись логов
f, err := os.Create("mylog.log")
gin.DefaultWriter = io.MultiWriter(f)

func TestXxx (* testing.T) - /*Пакетное тестирование обеспечивает поддержку автоматического тестирования пакетов Go.
Он предназначен для использования вместе с командой "go test", которая автоматизирует выполнение любой функции формы,
где Xxx не начинается с строчной буквы. Имя функции служит для идентификации процедуры тестирования.*/
