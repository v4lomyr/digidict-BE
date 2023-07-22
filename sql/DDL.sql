create table if not exists words (
	word_id BIGSERIAL primary key,
	word VARCHAR not null,
	word_class VARCHAR not null,
	word_description VARCHAR not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	example_phrase VARCHAR not null
);

create table if not exists indonesian_words (
	indonesian_word_id BIGSERIAL primary key,
	word_id BIGINT not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	foreign key (word_id) references words (word_id)
);

create table if not exists english_words (
	english_word_id BIGSERIAL primary key,
	word_id BIGINT not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	foreign key (word_id) references words (word_id)
);

create table if not exists languages (
	language_id SERIAL primary key,
	language VARCHAR not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp
);

create table if not exists translations (
	translation_id BIGSERIAL primary key,
	source_language_id INT not null,
	translation_language_id INT not null,
	source_word_id BIGINT not null,
	translation_word_id BIGINT not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	foreign key (source_language_id) references languages (language_id),
	foreign key (translation_language_id) references languages (language_id),
	foreign key (source_word_id) references words (word_id),
	foreign key (translation_word_id) references words (word_id)
);
