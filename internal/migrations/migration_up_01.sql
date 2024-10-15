CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid(),
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    picture text,
    email_verified boolean DEFAULT FALSE,

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
    CONSTRAINT FK_recruiter_user_id FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
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

CREATE TABLE talent_saved_jobs (
    talent_id uuid,
    job_id uuid,

    CONSTRAINT PK_saved_job PRIMARY KEY (talent_id, job_id),
    CONSTRAINT FK_saved_job_talent FOREIGN KEY (talent_id) REFERENCES "talents"(id),
    CONSTRAINT FK_saved_job_job FOREIGN KEY (job_id) REFERENCES "jobs"(id)
);