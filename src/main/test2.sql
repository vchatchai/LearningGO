WITH advance_payment AS 
    (SELECT f.con_id,
         COUNT(distinct(to_char( f.cfw_dt_installment,
         'YYYYMM'))) AS total_month_advance_paid
    FROM td_cashflow f
    WHERE f.cfw_typ_id IN (2, 5)
            AND f.cfw_bl_cancel IS FALSE
            AND f.cfw_bl_paid IS TRUE
            AND to_timestamp( ? / 1000) < f.cfw_dt_installment
            AND f.cfw_am_te_installment > 0
    GROUP BY  con_id ), last_payment AS 
    (SELECT f.con_id,
         max(pay.paymentdate) AS paymentdate
    FROM td_cashflow f
    INNER JOIN td_ar_payment ar
        ON f.cfw_va_payment_reference = ar.reference
    INNER JOIN td_ar_payment pay
        ON ar.reference = pay.reference
    WHERE f.cfw_va_payment_reference IS NOT NULL
            AND f.cfw_am_te_installment > 0
    GROUP BY  f.con_id ) , unpaid AS 
    (SELECT c.con_id,
        COUNT(distinct(to_char( f.cfw_dt_installment,
         'YYYYMM'))) AS total_month_unpaid
    FROM td_contract c
    INNER JOIN td_cashflow f
        ON c.con_id = f.con_id
    LEFT JOIN td_collection col
        ON c.con_id = col.con_id
    WHERE f.cfw_typ_id IN (2, 5)
            AND f.cfw_bl_cancel IS FALSE
            AND f.cfw_bl_paid IS FALSE
            AND to_timestamp( ? / 1000) > f.cfw_dt_installment
            AND f.cfw_am_te_installment > 0
    GROUP BY  c.con_id ), partial AS 
    (SELECT c.con_id,
         COUNT(distinct(f.cfw_bl_paid)) > 1 AS is_partial
    FROM td_contract c
    INNER JOIN td_cashflow f
        ON c.con_id = f.con_id
    WHERE f.cfw_bl_cancel IS FALSE
            AND to_char(f.cfw_dt_installment, 'YYYYMM') = to_char(to_timestamp(? / 1000), 'YYYYMM')
            AND f.cfw_am_te_installment > 0
    GROUP BY  c.con_id ), next_installment AS 
    (SELECT con.con_id,
         round(f.cfw_am_te_installment::numeric ,
        2 ) AS interest_next_installment,
         f.cfw_dt_installment AS next_installment_date
    FROM td_cashflow f
    INNER JOIN td_contract con
        ON con.con_id = f.con_id
    WHERE f.cfw_typ_id IN (2)
            AND to_char(f.cfw_dt_installment, 'YYYYMM') = to_char(to_timestamp(? / 1000) + INTERVAL '1 month', 'YYYYMM')
            AND con.wkf_sta_id IN (202)
            AND f.cfw_interest_realized is NULL
            AND con.interest_realized_uncalculated_date is NULL )
SELECT f.cfw_id,
         con.con_va_reference,
         round(f.cfw_am_te_installment::numeric ,
         2 ) AS interest,
         f.cfw_dt_installment,
         DATE_TRUNC('month', f.cfw_dt_installment) + INTERVAL '1 month' - INTERVAL '1 day' AS END_OF_MONTH, n.next_installment_date AS NEXT_INSTALLMENT, EXTRACT(DAY
FROM (DATE_TRUNC('month', f.cfw_dt_installment) + INTERVAL '1 month' ) - f.cfw_dt_installment ) AS REMAINING_DATE, EXTRACT(DAY
FROM (f.cfw_dt_installment + INTERVAL '1 month') - f.cfw_dt_installment ) PERIOD_OF_INSTALLMENT, con.wkf_sta_id, coalesce(total_month_unpaid,0) AS total_month_unpaid, coalesce(a.total_month_advance_paid,0) AS total_month_advance_paid, payment.paymentdate, con.con_dt_start AS contract_date, p.is_partial AS is_partial, f.cfw_bl_paid AS is_paid, n.interest_next_installment AS interest_next_installment
FROM td_cashflow f
INNER JOIN td_contract con
    ON con.con_id = f.con_id
INNER JOIN partial p
    ON p.con_id = con.con_id
INNER JOIN next_installment n
    ON n.con_id = con.con_id
LEFT JOIN unpaid u
    ON con.con_id = u.con_id
LEFT JOIN advance_payment a
    ON con.con_id = a.con_id
LEFT JOIN last_payment payment
    ON payment.con_id = con.con_id
WHERE f.cfw_typ_id IN ( 2, 9)
        AND to_char(f.cfw_dt_installment, 'YYYYMM') = to_char(to_timestamp(? / 1000), 'YYYYMM')
        AND con.wkf_sta_id IN (202)
        AND f.cfw_interest_realized is NULL
        AND con.interest_realized_uncalculated_date is NULL 