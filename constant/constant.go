package constant

const (
	ServiceDateWise = `
		select
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY') as datewise,
	service,
	SUM(patient_share_amount::float) as amount,
	CASE 
        WHEN RANK() OVER (PARTITION BY TO_CHAR(order_date::TIMESTAMP, 'MM/DD/YYYY') ORDER BY service) = 1
        THEN SUM(SUM(patient_share_amount::float)) 
             OVER (PARTITION BY TO_CHAR(order_date::TIMESTAMP, 'MM/DD/YYYY'))
        ELSE NULL
    END AS daily_total
from
	bill_masters bm
group by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'),
	service
order by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'), service
	`
	DoctorServiceSpecific = `
	select
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY') as datewise,
	service,
	SUM(patient_share_amount::float) as amount,
    doctor_name,
    CASE 
        WHEN RANK() OVER (PARTITION BY TO_CHAR(order_date::TIMESTAMP, 'MM/DD/YYYY') ORDER BY service) = 1
        THEN SUM(SUM(patient_share_amount::float)) 
             OVER (PARTITION BY TO_CHAR(order_date::TIMESTAMP, 'MM/DD/YYYY'))
        ELSE NULL
    END AS daily_total
from
	bill_masters bm
group by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'),
	service,
	doctor_name 
order by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'), service, doctor_name;
	`
	Summary = `
	select 
	service,
	SUM(patient_share_amount::float) as amount
from 
	bill_masters bm 
group by 
	service;
	`
	ItemWiseSpecification = `
	select
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY') as datewise,
	service,
	oracle_code,
	item,
	count(quantity) as quantity,
	SUM(patient_share_amount::float) as amount
from
	bill_masters bm
group by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'),
	service,
	oracle_code,
	item
order by
	TO_CHAR(order_date::TIMESTAMP,
	'MM/DD/YYYY'), service,oracle_code, item;
	`
)
