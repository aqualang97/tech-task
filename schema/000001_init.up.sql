create table if not exists data
(
    transaction_id      integer          not null        unique,
    request_id          integer          not null,
    terminal_id         integer          not null,
    partner_object_id   integer          not null,
    amount_total        integer          not null,
    amount_original     integer          not null,
    commission_ps       double precision not null,
    commission_client   integer          not null,
    commission_provider double precision not null,
    date_input          timestamp        not null,
    date_post           timestamp        not null,
    status              varchar(10)      not null,
    payment_type        varchar(10)      not null,
    payment_number      varchar(15)      not null,
    service_id          integer          not null,
    service             varchar(20)      not null,
    payee_id            integer          not null,
    payee_name          varchar(10)      not null,
    payee_bank_mfo      integer          not null,
    payee_bank_account  text             not null,
    payment_narrative   text             not null
);
