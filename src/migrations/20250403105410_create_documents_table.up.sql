CREATE TABLE file_records (
    id UUID PRIMARY KEY,
    file_name TEXT,
    provider TEXT,
    url TEXT,
    object_key TEXT,
    uploaded_at TIMESTAMP
    description TEXT,
);