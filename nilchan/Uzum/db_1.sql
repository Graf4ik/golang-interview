-- Все индефикаторы магазинов и кол-во заказов в них,
-- в которых сделано больше 100 заказов за сентябрь 23-го

CREATE TABLE orders(
    id INTEGER PRIMARY KEY NOT NULL,
    sum INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    shop_id INTEGER NOT NULL,
    created_at timestamp not null,
    updated_at timestamp not null,
)

SELECT id, COUNT(id) as orders_count
FROM orders
WHERE created_at >= "2023-01-09" AND created_at <= "2023-10-01"
GROUP BY shop_id
HAVING COUNT(id) > 100

