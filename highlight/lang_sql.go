package highlight

// SQL według ISO/IEC 9075:2016 (słowa zastrzeżone, literały, komentarze) oraz słów powszechnych w dokumentacji
// PostgreSQL, SQLite i MySQL (LIMIT, INDEX, IF, AUTO_INCREMENT); wymagania: requirements/languages/sql.md.
var langSQL = language{
	names: []string{"sql"},
	scan: (&table{
		words: foldWords(words(
			wordGroup{keyword, `ADD ALL ALTER AND ANY AS ASC AUTO_INCREMENT BEGIN BETWEEN BY CASCADE CASE CAST
				CHECK COLLATE COLUMN COMMIT CONSTRAINT CREATE CROSS CURRENT_DATE CURRENT_TIME CURRENT_TIMESTAMP
				CURRENT_USER DECLARE DEFAULT DELETE DESC DISTINCT DROP ELSE END ESCAPE EXCEPT EXISTS FETCH
				FIRST FOR FOREIGN FROM FULL FUNCTION GRANT GROUP HAVING IF IN INDEX INNER INSERT INTERSECT INTO
				IS JOIN KEY LEFT LIKE LIMIT NATURAL NEXT NOT OFFSET ON ONLY OR ORDER OUTER OVER PARTITION
				PRIMARY PROCEDURE RECURSIVE REFERENCES RETURN RETURNS REVOKE RIGHT ROLLBACK ROW ROWS SELECT SET
				SOME TABLE THEN TO TRANSACTION TRIGGER UNION UNIQUE UPDATE USING VALUES VIEW WHEN WHERE WINDOW
				WITH`},
			wordGroup{typeName, `BIGINT BINARY BLOB BOOLEAN CHAR CHARACTER CLOB DATE DECIMAL DOUBLE FLOAT INT
				INTEGER INTERVAL NUMERIC PRECISION REAL SMALLINT TEXT TIME TIMESTAMP VARBINARY VARCHAR VARYING`},
			wordGroup{builtin, `TRUE FALSE NULL UNKNOWN ABS AVG CHAR_LENGTH COALESCE COUNT LOWER MAX MIN NULLIF
				SUBSTRING SUM TRIM UPPER`},
		)),
		foldCase:      true,
		lineComments:  []string{"--"},
		blockComments: []blockComment{{open: "/*", close: "*/", nested: true}},
		quotes: []quote{
			{open: `'`, close: `'`, doubled: true, multiline: true},
			{open: `"`, close: `"`, doubled: true, ident: true},
			{open: "`", close: "`", doubled: true, ident: true},
		},
		prefixes: []prefix{{text: "N"}, {text: "n"}, {text: "X"}, {text: "x"}, {text: "B"}, {text: "b"},
			{text: "E"}, {text: "e"}, {text: "U&"}, {text: "u&"}},
		operators: "+-*/%<>=!|&^~:",
	}).scan,
}
