import { MigrationInterface, QueryRunner } from "typeorm";

export class InitialSchema1738361203895 implements MigrationInterface {
    name = 'InitialSchema1738361203895'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`CREATE TYPE "public"."belt_progress_belt_enum" AS ENUM('WHITE', 'BLUE', 'PURPLE', 'BROWN', 'BLACK')`);
        await queryRunner.query(`CREATE TABLE "belt_progress" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "belt" "public"."belt_progress_belt_enum" NOT NULL, "stripes" integer NOT NULL DEFAULT '0', "achievedAt" TIMESTAMP NOT NULL DEFAULT now(), "userId" uuid, CONSTRAINT "PK_a79b99afa7f9b883e5a50c5f6a5" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE TYPE "public"."technique_type_enum" AS ENUM('SUBMISSION', 'SWEEP', 'TAKEDOWN', 'GUARD')`);
        await queryRunner.query(`CREATE TYPE "public"."technique_difficulty_enum" AS ENUM('BASIC', 'INTERMEDIATE', 'ADVANCED')`);
        await queryRunner.query(`CREATE TABLE "technique" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "name" character varying NOT NULL, "description" character varying NOT NULL, "type" "public"."technique_type_enum" NOT NULL, "difficulty" "public"."technique_difficulty_enum" NOT NULL DEFAULT 'BASIC', CONSTRAINT "UQ_f61ae45d7238c195698df9e8c62" UNIQUE ("name"), CONSTRAINT "PK_a149a5ec01b7bf13dbeb2f743b5" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE TABLE "training_session" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "date" TIMESTAMP NOT NULL, "duration" integer NOT NULL, "note" character varying NOT NULL, "userId" uuid, CONSTRAINT "PK_a17a9657ff5a6e048bfd82c4651" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE TABLE "user" ("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "name" character varying NOT NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "createdAt" TIMESTAMP NOT NULL DEFAULT now(), CONSTRAINT "UQ_e12875dfb3b1d92d7d7c5377e22" UNIQUE ("email"), CONSTRAINT "PK_cace4a159ff9f2512dd42373760" PRIMARY KEY ("id"))`);
        await queryRunner.query(`CREATE TABLE "training_session_techniques_technique" ("trainingSessionId" uuid NOT NULL, "techniqueId" uuid NOT NULL, CONSTRAINT "PK_c6b30054b31ddc9ff6dbfbd911a" PRIMARY KEY ("trainingSessionId", "techniqueId"))`);
        await queryRunner.query(`CREATE INDEX "IDX_86e9c59d8e988451fc79f351ef" ON "training_session_techniques_technique" ("trainingSessionId") `);
        await queryRunner.query(`CREATE INDEX "IDX_cea7191135ce891b0c6240c721" ON "training_session_techniques_technique" ("techniqueId") `);
        await queryRunner.query(`ALTER TABLE "belt_progress" ADD CONSTRAINT "FK_9ac3ea424c67a3f0c0b3894ca3e" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE NO ACTION`);
        await queryRunner.query(`ALTER TABLE "training_session" ADD CONSTRAINT "FK_487b6d452df5807077415b6d080" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE NO ACTION`);
        await queryRunner.query(`ALTER TABLE "training_session_techniques_technique" ADD CONSTRAINT "FK_86e9c59d8e988451fc79f351ef3" FOREIGN KEY ("trainingSessionId") REFERENCES "training_session"("id") ON DELETE CASCADE ON UPDATE CASCADE`);
        await queryRunner.query(`ALTER TABLE "training_session_techniques_technique" ADD CONSTRAINT "FK_cea7191135ce891b0c6240c721d" FOREIGN KEY ("techniqueId") REFERENCES "technique"("id") ON DELETE CASCADE ON UPDATE CASCADE`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "training_session_techniques_technique" DROP CONSTRAINT "FK_cea7191135ce891b0c6240c721d"`);
        await queryRunner.query(`ALTER TABLE "training_session_techniques_technique" DROP CONSTRAINT "FK_86e9c59d8e988451fc79f351ef3"`);
        await queryRunner.query(`ALTER TABLE "training_session" DROP CONSTRAINT "FK_487b6d452df5807077415b6d080"`);
        await queryRunner.query(`ALTER TABLE "belt_progress" DROP CONSTRAINT "FK_9ac3ea424c67a3f0c0b3894ca3e"`);
        await queryRunner.query(`DROP INDEX "public"."IDX_cea7191135ce891b0c6240c721"`);
        await queryRunner.query(`DROP INDEX "public"."IDX_86e9c59d8e988451fc79f351ef"`);
        await queryRunner.query(`DROP TABLE "training_session_techniques_technique"`);
        await queryRunner.query(`DROP TABLE "user"`);
        await queryRunner.query(`DROP TABLE "training_session"`);
        await queryRunner.query(`DROP TABLE "technique"`);
        await queryRunner.query(`DROP TYPE "public"."technique_difficulty_enum"`);
        await queryRunner.query(`DROP TYPE "public"."technique_type_enum"`);
        await queryRunner.query(`DROP TABLE "belt_progress"`);
        await queryRunner.query(`DROP TYPE "public"."belt_progress_belt_enum"`);
    }

}
