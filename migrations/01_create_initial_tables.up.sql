DROP TABLE IF EXISTS shift_schedule CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;

CREATE TABLE IF NOT EXISTS shift_schedule (
 	id SERIAL PRIMARY KEY,
    alias VARCHAR(255) NOT NULL,
    organization JSONB NOT NULL, -- group name, mail, phone, description
    manager JSONB NOT NULL,
    description VARCHAR(1024) DEFAULT NULL, 
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    shifts JSONB DEFAULT NULL,
    year INTEGER NOT NULL,
    status INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Shift Schedule 1
INSERT INTO shift_schedule (
    alias,
    organization,
    manager,
    description,
    start_date,
    end_date,
    shifts,
    year,
    status
) VALUES (
    'Shift 1',
    '[{"id": 1, "name": "Group 1", "mail": "", "phone": "", "description": ""}]',
    '[{"id": 1, "name": "Manager 1", "mail": "", "phone": "", "description": ""}]',
    'Description 1',
    '2023-01-01 00:00:00',
    '2023-10-01 00:00:00',
    '[
        {
            "id": 0,
            "user": {
                "id": 21304362,
                "name": "User 1",
                "mail": "",
                "phone": "",
                "description": ""

            },
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00"
        },
        {
            "id": 1,
            "user": {
                "id": 21304362,
                "name": "User 2",
                "mail": "",
                "phone": "",
                "description": ""

            },
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00"
        },
        {
            "id": 2,
            "user": {
                "id": 21304362,
                "name": "User 3",
                "mail": "",
                "phone": "",
                "description": ""

            },
            "start": "2023-01-01 00:00:00",
            "end": "2023-02-01 00:00:00"
        }
    ]',
    2023,
    0
);