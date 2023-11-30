CREATE TABLE "session" (
  "id" serial PRIMARY KEY NOT NULL,
  "startTime" timestamp NOT NULL,
  "endTime" timestamp,
  "duration" interval,
  "tags" varchar(20)[]
);

CREATE INDEX ON "session" ("duration");

CREATE INDEX ON "session" ("tags");
