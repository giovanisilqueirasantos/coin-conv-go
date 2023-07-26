CREATE TABLE coinconv.exchanges (
	currency_from varchar(100) NOT NULL,
	currency_to varchar(100) NOT NULL,
	amount varchar(100) NOT NULL,
	rate varchar(100) NOT NULL,
	conv_amount varchar(100) NOT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;
