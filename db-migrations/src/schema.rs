table! {
    categories (id) {
        id -> Int4,
        uuid -> Uuid,
        title -> Varchar,
        description -> Nullable<Varchar>,
        created_at -> Timestamp,
        updated_at -> Nullable<Timestamp>,
    }
}

table! {
    tb_categories (id) {
        id -> Int4,
        uuid -> Uuid,
        title -> Varchar,
        description -> Nullable<Varchar>,
        created_at -> Timestamp,
        updated_at -> Nullable<Timestamp>,
    }
}

allow_tables_to_appear_in_same_query!(
    categories,
    tb_categories,
);
