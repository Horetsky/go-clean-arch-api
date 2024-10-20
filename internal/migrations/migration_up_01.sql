CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE PROFILE_TYPE AS ENUM ('TALENT', 'RECRUITER');

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(),
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    picture text,
    email_verified boolean DEFAULT FALSE,
    type profile_type

    CONSTRAINT PK_user PRIMARY KEY (id)
);

CREATE TABLE talents (
    id uuid DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL UNIQUE,

    CONSTRAINT PK_talent PRIMARY KEY (id),
    CONSTRAINT FK_talent_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE talent_profiles (
    id uuid DEFAULT gen_random_uuid(),
    talent_id uuid NOT NULL,
    category VARCHAR(100) NOT NULL,

    first_name varchar(100),
    last_name varchar(100),

    phone varchar(100),
    linkedIn_url text,
    resume_url text,
    photo text,

    CONSTRAINT PK_talent_profile PRIMARY KEY (id),
    CONSTRAINT FK_talent_profile_talent FOREIGN KEY (talent_id) REFERENCES talents(id) ON DELETE CASCADE
);

CREATE TABLE recruiters (
    id uuid DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL UNIQUE,

    CONSTRAINT PK_recruiter_id PRIMARY KEY (id),
    CONSTRAINT FK_recruiter_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE recruiter_profiles (
    id uuid DEFAULT gen_random_uuid(),
    recruiter_id uuid NOT NULL,

    first_name varchar(100),
    last_name varchar(100),
    phone varchar(100),
    linkedIn_url text,
    company_name varchar(100),
    company_website_url text,

    CONSTRAINT PK_recruiter_profile PRIMARY KEY (id),
    CONSTRAINT FK_recruiter_profile_recruiter FOREIGN KEY (recruiter_id) REFERENCES recruiters(id) ON DELETE CASCADE
);

CREATE TABLE jobs (
    id uuid DEFAULT gen_random_uuid(),
    recruiter_id uuid NOT NULL,
    category VARCHAR(100) NOT NULL,

    title text NOT NULL,
    description text NOT NULL,
    requirements text NOT NULL,

    CONSTRAINT PK_job PRIMARY KEY (id),
    CONSTRAINT FK_job_recruiter FOREIGN KEY (recruiter_id) REFERENCES recruiters(id) ON DELETE CASCADE
);

CREATE table job_applications (
    talent_id uuid NOT NULL,
    job_id uuid NOT NULL,

    CONSTRAINT U_talent_job UNIQUE (talent_id, job_id),
    CONSTRAINT FK_application_talent FOREIGN KEY (talent_id) REFERENCES talents(id),
    CONSTRAINT FK_application_job FOREIGN KEY (job_id) REFERENCES jobs(id)
);