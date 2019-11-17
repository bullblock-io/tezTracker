CREATE INDEX ix_opeations_transactions_source_operation_id
    ON tezos.operations USING btree
    (source, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_opeations_transactions_destination_operation_id
    ON tezos.operations USING btree
    (destination, operation_id)
      WHERE kind='transaction';

CREATE INDEX ix_opeations_endorsements_delegate_block_level
    ON tezos.operations USING btree
    (delegate, block_level)
      WHERE kind='endorsement';

CREATE INDEX ix_opeations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_opeations_delegations_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='delegation';

CREATE INDEX ix_opeations_endorsements_operation_id
    ON tezos.operations USING btree
    (operation_id)
      WHERE kind='endorsement';

CREATE INDEX ix_accounts_balance
    ON tezos.accounts USING btree
    (balance);



Create table tezos.operation_counters(
    cnt_id SERIAL PRIMARY KEY,
    cnt_last_op_id int not null,
    cnt_operation_type varchar(100) NOT NULL,
    cnt_count bigint not null,
    cnt_created_at timestamp with time zone NULL DEFAULT NULL,
    CONSTRAINT operation_counters_last_op_foreign FOREIGN KEY (cnt_last_op_id) REFERENCES tezos.operations (operation_id)
);


CREATE TABLE tezos.future_baking_rights (
    level integer NOT NULL,
    delegate character varying NOT NULL,
    priority integer NOT NULL,
    estimated_time timestamp without time zone NOT NULL,
    PRIMARY KEY (level, priority)
);

CREATE INDEX future_baking_rights_delegate_idx 
    ON tezos.future_baking_rights USING btree (delegate);

alter table tezos.blocks ADD UNIQUE (level);

CREATE TABLE tezos.snapshots (
    snp_cycle integer PRIMARY KEY ,
    snp_block_level integer NOT NULL,
    snp_rolls integer NOT NULL,
    CONSTRAINT snapshots_block_foreign FOREIGN KEY (snp_block_level) REFERENCES tezos.blocks (level)
);
