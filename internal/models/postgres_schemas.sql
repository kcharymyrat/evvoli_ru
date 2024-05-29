CREATE TABLE language (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE category (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    image TEXT NOT NULL,
    thumbnail TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    modified_at TIMESTAMP WITH TIME ZONE,
    -- created_by_id UUID REFERENCES auth_user(id) ON DELETE CASCADE
);

CREATE TABLE category_translation (
    id BIGSERIAL PRIMARY KEY,
    category_id UUID NOT NULL REFERENCES category(id) ON DELETE CASCADE,
    language_id INTEGER NOT NULL REFERENCES language(id),
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    UNIQUE (category_id, language_id)
);


-- Create a function to set the created_at and modified_at fields
CREATE OR REPLACE FUNCTION set_timestamp() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.created_at IS NULL THEN
        NEW.created_at := CURRENT_TIMESTAMP;
    END IF;
    NEW.modified_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to set the timestamps on insert and update
CREATE TRIGGER set_timestamp_on_insert
BEFORE INSERT ON category
FOR EACH ROW EXECUTE FUNCTION set_timestamp();

CREATE TRIGGER set_timestamp_on_update
BEFORE UPDATE ON category
FOR EACH ROW EXECUTE FUNCTION set_timestamp();


CREATE OR REPLACE FUNCTION generate_slug(input_text TEXT)
RETURNS TEXT AS $$
DECLARE
    slug_text TEXT;
BEGIN
    -- Remove special characters and diacritics, convert to lowercase
    slug_text := lower(regexp_replace(input_text, '[^\w\s]', '', 'g'));

    -- Replace spaces with dashes
    slug_text := regexp_replace(slug_text, '\s+', '-', 'g');

    -- Trim leading and trailing dashes
    slug_text := regexp_replace(slug_text, '^[-]+|[-]+$', '', 'g');

    RETURN slug_text;
END;
$$ LANGUAGE plpgsql;
