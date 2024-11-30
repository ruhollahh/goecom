INSERT INTO admins (id, name, phone_number, hashed_password, active, created_at, updated_at) VALUES
    ('5cf37266-3473-4006-984f-9325122678b7', 'Admin Gopher', '09384664416', '$2a$10$qVjYHVYolxXTeTYD2pwzHukbffR/heH8m9QdAsP92U7Moi2Pub1hm', true, NOW(), NOW())
ON CONFLICT DO NOTHING;
