Необходимо реализовать конкурентный поиск документов на серверах. Для этого у нас есть сторонняя библиотека с функцией
search.Search(server string, query string) ([]string, error)
которая осуществляет поиск документов на указанном сервере по указанному запросу. У нас есть N серверов идентичных (реплики) и задача состоит в том, чтобы конкурентно вызвать эту функцию для всех серверов и вернуть первый успешный ответ от любого из серверов не дожидаясь ответов от других серверов. Если какой-то сервер возвращает ошибку, то мы ее игнорируем, дожидаясь успешного ответа от других, но если все сервера ответили ошибкой, то наша функция должна вернуть ошибку, что поиск не удался.
func Search(servers []string, query string) ([]string, error) {