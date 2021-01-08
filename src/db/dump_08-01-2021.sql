--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4 (Debian 12.4-1.pgdg100+1)
-- Dumped by pg_dump version 12.4 (Debian 12.4-1.pgdg100+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE ONLY benefits.related_benefits DROP CONSTRAINT related_benefit_id_fk;
ALTER TABLE ONLY benefits.life_journey_benefit DROP CONSTRAINT life_journey_fk;
ALTER TABLE ONLY benefits.related_benefits DROP CONSTRAINT benefit_id_fk;
ALTER TABLE ONLY benefits.life_journey_benefit DROP CONSTRAINT benefit_fk;
ALTER TABLE ONLY benefits.related_benefits DROP CONSTRAINT related_benefits_pkey;
ALTER TABLE ONLY benefits.life_journey DROP CONSTRAINT life_journey_pkey;
ALTER TABLE ONLY benefits.life_journey_benefit DROP CONSTRAINT life_journey_benefit_pkey;
ALTER TABLE ONLY benefits.benefit DROP CONSTRAINT benefit_pkey;
DROP TABLE benefits.related_benefits;
DROP TABLE benefits.life_journey_benefit;
DROP TABLE benefits.life_journey;
DROP TABLE benefits.benefit;
DROP SCHEMA benefits;
--
-- Name: benefits; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA benefits;


ALTER SCHEMA benefits OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: benefit; Type: TABLE; Schema: benefits; Owner: postgres
--

CREATE TABLE benefits.benefit (
    id character varying(8) NOT NULL,
    title character varying(32) NOT NULL,
    description character varying(64),
    long_description character varying(2048)
);


ALTER TABLE benefits.benefit OWNER TO postgres;

--
-- Name: life_journey; Type: TABLE; Schema: benefits; Owner: postgres
--

CREATE TABLE benefits.life_journey (
    id character varying(8) NOT NULL,
    title character varying(32) NOT NULL,
    description character varying(2048)
);


ALTER TABLE benefits.life_journey OWNER TO postgres;

--
-- Name: life_journey_benefit; Type: TABLE; Schema: benefits; Owner: postgres
--

CREATE TABLE benefits.life_journey_benefit (
    life_journey_id character varying(8) NOT NULL,
    benefit_id character varying(8) NOT NULL
);


ALTER TABLE benefits.life_journey_benefit OWNER TO postgres;

--
-- Name: related_benefits; Type: TABLE; Schema: benefits; Owner: postgres
--

CREATE TABLE benefits.related_benefits (
    benefit_id character varying(8) NOT NULL,
    related_benefit_id character varying(8) NOT NULL
);


ALTER TABLE benefits.related_benefits OWNER TO postgres;

--
-- Data for Name: benefit; Type: TABLE DATA; Schema: benefits; Owner: postgres
--

COPY benefits.benefit (id, title, description, long_description) FROM stdin;
1	Driver's license	License for Driver's	# Driver's License Benefit Details \\nBrief description of the Driver's License Benefit \\n## Overview \\nShort Overview of Driver's License Process \\n## Important Information \\nImportant Information cli8ents need to know for the Driver's License benefit. \\n## Eligibility criteria \\nDescription of Driver's License Eligibility Criteria and how to qualify for the benefit. \\n- Must be 16 years of age or older \\n- Must be a resident of Canada \\n- Must have completed a Driver Training Program \\n- For client's under the age of 18, you must have parental consent \\n### Examples \\nDescription of different scenario's to provide examples to clients \\n## Eligibility period \\nDescription of Eligibility periods for a Driver's License. \\n## How to apply \\nDescription of How to apply for a Driver's License and what information is required. \\n## Contact Information \\nFor further Information on Driver's License and related Benefits contact 1-800-Drivers. \\n## Payment Information \\nDescription of Payment Information for a Driver's License.
2	Student Card	Student Card Benefit	# Student Card Benefit Details \\nBrief description of the Student Card Benefit \\n## Overview \\nShort Overview of Student Card Process \\n## Important Information \\nImportant Information clients need to know for the Student Card benefit. \\n## Eligibility criteria \\nDescription of Student Card Eligibility Criteria and how to qualify for the benefit. \\n- Must be under 18 years of age \\n- Must be a resident of Canada \\n- Must have parental consent \\n### Examples \\nDescription of different scenario's to provide examples to clients \\n## Eligibility period \\nDescription of Eligibility periods for a Student Card. \\n## How to apply \\nDescription of How to apply for a Student Card and what information is required. \\n## Contact Information \\nFor further Information on Student Card and related Benefits contact 1-800-student. \\n## Payment Information \\nDescription of Payment Information for a Student Card.
3	Senior's Card	Senior card benefit	# Senior card Benefit Details \\nBrief description of the Senior card Benefit \\n## Overview \\nShort Overview of Senior card Process \\n## Important Information \\nImportant Information cli8ents need to know for the Senior card benefit. \\n## Eligibility criteria \\nDescription of Senior card Eligibility Criteria and how to qualify for the benefit. \\n- Must be over 55 years of age \\n- Must be a resident of Canada \\n- Must be retired \\n### Examples \\nDescription of different scenario's to provide examples to clients \\n## Eligibility period \\nDescription of Eligibility periods for a Senior card. \\n## How to apply \\nDescription of How to apply for a Senior card and what information is required. \\n## Contact Information \\nFor further Information on Senior card and related Benefits contact 1-800-seniors. \\n## Payment Information \\nDescription of Payment Information for a Senior card.
4	Disability Card	Disability card is a benefit for peoples with a disability	# Disability card Benefit Details \\nBrief description of the Disability card Benefit \\n## Overview \\nShort Overview of Disability card Process \\n## Important Information \\nImportant Information clients need to know for the Disability card benefit. \\n## Eligibility criteria \\nDescription of Disability card Eligibility Criteria and how to qualify for the benefit. \\n- Must have a valid medical Card \\n- Must be a resident of Canada \\n- Must have a medical Certificate \\n### Examples \\nDescription of different scenario's to provide examples to clients \\n## Eligibility period \\nDescription of Eligibility periods for a Disability card. \\n## How to apply \\nDescription of How to apply for a Disability card and what information is required. \\n## Contact Information \\nFor further Information on Disability card and related Benefits contact 1-800-OCANADA. \\n## Payment Information \\nDescription of Payment Information for a Disability card.
5	Medical Insurance Card	Medical Insurance card is a benefit for Canadian's	# Medical Insurance card Benefit Details \\nBrief description of the Medical Insurance card Benefit \\n## Overview \\nShort Overview of Medical Insurance card Process \\n## Important Information \\nImportant Information clients need to know for the Medical Insurance card benefit. \\n## Eligibility criteria \\nDescription of Medical Insurance card Eligibility Criteria and how to qualify for the benefit. \\n- Must have a permanant address \\n- Must be a resident of Canada \\n- Must be a Canadian Citizen \\n### Examples \\nDescription of different scenario's to provide examples to clients \\n## Eligibility period \\nDescription of Eligibility periods for a Medical Insurance card. \\n## How to apply \\nDescription of How to apply for a Medical Insurance card and what information is required. \\n## Contact Information \\nFor further Information on Medical Insurance card and related Benefits contact 1-800-OCANADA. \\n## Payment Information \\nDescription of Payment Information for a Medical Insurance card.
\.


--
-- Data for Name: life_journey; Type: TABLE DATA; Schema: benefits; Owner: postgres
--

COPY benefits.life_journey (id, title, description) FROM stdin;
1	Adult Life Journey	Adult Life Journey relates all benefits to support Adults
2	Student Life Journey	Student Life Journey relates all benefits to support Students
3	Senior Life Journey	Senior Life Journey relates all benefits to support Seniors
4	Disability Life Journey	Disability Life Journey relates all benefits to support disabled people
\.


--
-- Data for Name: life_journey_benefit; Type: TABLE DATA; Schema: benefits; Owner: postgres
--

COPY benefits.life_journey_benefit (life_journey_id, benefit_id) FROM stdin;
1	1
1	4
1	5
2	1
2	2
2	4
2	5
3	1
3	3
3	4
3	5
4	4
4	5
\.


--
-- Data for Name: related_benefits; Type: TABLE DATA; Schema: benefits; Owner: postgres
--

COPY benefits.related_benefits (benefit_id, related_benefit_id) FROM stdin;
1	5
2	1
2	5
3	1
3	4
3	5
4	5
5	4
\.


--
-- Name: benefit benefit_pkey; Type: CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.benefit
    ADD CONSTRAINT benefit_pkey PRIMARY KEY (id);


--
-- Name: life_journey_benefit life_journey_benefit_pkey; Type: CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.life_journey_benefit
    ADD CONSTRAINT life_journey_benefit_pkey PRIMARY KEY (life_journey_id, benefit_id);


--
-- Name: life_journey life_journey_pkey; Type: CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.life_journey
    ADD CONSTRAINT life_journey_pkey PRIMARY KEY (id);


--
-- Name: related_benefits related_benefits_pkey; Type: CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.related_benefits
    ADD CONSTRAINT related_benefits_pkey PRIMARY KEY (benefit_id, related_benefit_id);


--
-- Name: life_journey_benefit benefit_fk; Type: FK CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.life_journey_benefit
    ADD CONSTRAINT benefit_fk FOREIGN KEY (benefit_id) REFERENCES benefits.benefit(id);


--
-- Name: related_benefits benefit_id_fk; Type: FK CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.related_benefits
    ADD CONSTRAINT benefit_id_fk FOREIGN KEY (benefit_id) REFERENCES benefits.benefit(id);


--
-- Name: life_journey_benefit life_journey_fk; Type: FK CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.life_journey_benefit
    ADD CONSTRAINT life_journey_fk FOREIGN KEY (life_journey_id) REFERENCES benefits.life_journey(id);


--
-- Name: related_benefits related_benefit_id_fk; Type: FK CONSTRAINT; Schema: benefits; Owner: postgres
--

ALTER TABLE ONLY benefits.related_benefits
    ADD CONSTRAINT related_benefit_id_fk FOREIGN KEY (related_benefit_id) REFERENCES benefits.benefit(id);


--
-- PostgreSQL database dump complete
--

