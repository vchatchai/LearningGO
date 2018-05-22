WITH X AS 
    (SELECT f.con_id,
         MAX(t.termination_letter_created ) AS cfw_create_termination_letter,
         count( DISTINCT f.cfw_nu_num_installment ) AS total_cfw_dt_installment,
         MIN(f.cfw_nu_num_installment ) AS cfw_nu_num_installment
    FROM td_contract c
    INNER JOIN td_cashflow f
        ON c.con_id = f.con_id
    INNER JOIN td_collection col
        ON c.con_id = col.con_id
    LEFT JOIN td_termination_letter_tracking t
        ON f.con_id = t.con_id
            AND t.cfw_nu_num_installment = f.cfw_nu_num_installment
    WHERE c.wkf_sta_id NOT IN (200, 213, 203, 500, 208, 206, 207)
            AND f.cfw_typ_id IN (5,2)
            AND f.cfw_bl_cancel IS FALSE
            AND f.cfw_bl_paid IS FALSE
            AND to_timestamp( 1526868592236 / 1000) > f.cfw_dt_installment
    GROUP BY  f.con_id
    ORDER BY  f.con_id )
SELECT con_id,
         cfw_create_termination_letter
FROM X
WHERE cfw_create_termination_letter IS NULL
        AND total_cfw_dt_installment > 2 