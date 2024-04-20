CREATE TABLE fact_status (
    company_id INT NOT NULL,
    car_id INT NOT NULL,
    date_id INT NOT NULL,
    PRIMARY KEY (company_id, car_id, driver_id, date_id)
);

SELECT
    dd.date,
    COALESCE(license_plate, 'N/A') as license_plate,
    COALESCE(trunc(extract(epoch from (MAX(dt.time) - MIN(dt.time))) / 3600.00, 2), 0.00)
FROM facts
    JOIN dim_company dc ON facts.company_id = dc.company_id
    JOIN (
        SELECT dim_car.*
        FROM dim_car
        WHERE car_key = (
            SELECT car_key FROM dim_car WHERE license_plate = 'G0PHER'
        )
    ) AS dcar ON facts.car_id = dcar.car_id
    RIGHT JOIN (SELECT * FROM dim_date WHERE month IN (3, 4)) AS dd ON facts.date_id = dd.date_id
GROUP BY dd.date, dcar.license_plate,
ORDER BY dd.date;
