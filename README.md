## clickhouse-admin-sms-consumer

### Проект реализует сервис по записи sms задач в clickhouse .

### Как работает сервис.

Данные для записи в clickhouse поступают по nats: 
 - subject - clickhouse:admin_sms_tasks

