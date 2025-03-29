-- Добавляем поля для данных авто
ALTER TABLE slots ADD COLUMN car_model TEXT;
ALTER TABLE slots ADD COLUMN car_number TEXT;
ALTER TABLE slots ADD COLUMN user_id INTEGER;

-- Создаем индекс для быстрого поиска записей пользователя
CREATE INDEX IF NOT EXISTS idx_slots_user_id ON slots(user_id);