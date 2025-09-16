-- Сделать рефакторинг запроса
select sender_code,
       receiver_code,
       content,
       document_schema_name,
       document_type,
       documents.doc_uuid as doc_uuid,
       permanent_object_uuit,
       create_Date,
       version,
       previous_container,
       sent_by_bus_at,
       invalid_signatures,
       documents.payload as payload,
       event.event as event
from documents
left join (
    select distinct on (doc_uuid) doc_uuid, event
    from doc_event --2млн
    order by doc_uuid, created_at desc
    ) as events on documents.doc_uuid = events.doc_uuid
    where true
        and lower(document_scheme_name) like 'aosr%'
        and permanent_object_uuid = 'd4d43f21-ba73-4e51-8dc2-edba2a6c9667'
        and events.event = ANY('{"ACCEPTED_BY_RECEIVER"}')
)
    order by create_date desc
limit 100
offset 0;

-- ✅ Отрефакторированный запрос:
-- WITH ... AS — это Common Table Expression (CTE) в SQL. Это способ задать временное именованное подмножество данных,
-- которое можно использовать в основном SQL-запросе, как если бы это была обычная таблица или представление.
WITH latest_events AS (
    SELECT DISTINCT ON (doc_uuid)
        doc_uuid,
        event
    FROM doc_event
    ORDER BY doc_uuid, created_at DESC
)

SELECT
    d.sender_code,
    d.receiver_code,
    d.content,
    d.document_schema_name,
    d.document_type,
    d.doc_uuid,
    d.permanent_object_uuid,
    d.create_date,
    d.version,
    d.previous_container,
    d.sent_by_bus_at,
    d.invalid_signatures,
    d.payload,
    e.event
FROM documents d
LEFT JOIN latest_events e ON d.doc_uuid = e.doc_uuid
WHERE
    lower(d.document_schema_name) LIKE 'aosr%' AND
    d.permanent_object_uuid = 'd4d43f21-ba73-4e51-8dc2-edba2a6c9667' AND
    e.event = ANY(ARRAY['ACCEPTED_BY_RECEIVER'])
ORDER BY d.create_date DESC
    LIMIT 100 OFFSET 0;

-- ⚠️ Возможные узкие места:
-- Таблица doc_event (2 млн записей) → DISTINCT ON + ORDER BY может быть медленным. Убедись, что есть индекс:
CREATE INDEX idx_doc_event_doc_uuid_created_at
ON doc_event (doc_uuid, created_at DESC);

-- 🔍 Что изменено и почему:
-- Изменение	                                Обоснование
-- WITH latest_events AS (...)	                Использование CTE (Common Table Expression) делает структуру запроса более читаемой.
-- DISTINCT ON (doc_uuid) → перенесён в CTE	    Это позволяет чётко отделить логику выбора последних событий.
-- Префикс d. и e.	                            Использование алиасов повышает читаемость и избегает коллизий.
-- ARRAY[...] вместо ANY('{...}')	            Явное использование массива понятнее и безопаснее.
-- where true and удалено	                    Избыточная конструкция.
-- Исправлена опечатка
-- document_scheme_name → document_schema_name  Возможно, в оригинале это была ошибка. Проверь.