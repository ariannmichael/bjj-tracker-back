import { TrainingSession } from "src/modules/training-session/domain/training-session.entity";
import { Column, Entity, ManyToMany, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class Technique {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ unique: true })
  name: string;

  @Column({ nullable: false })
  description: string;

  @Column({ type: 'enum', enum: ['SUBMISSION', 'SWEEP', 'TAKEDOWN', 'GUARD'] })
  type: string;

  @Column({ type: 'enum', enum: ['BASIC', 'INTERMEDIATE', 'ADVANCED'], default: 'BASIC' })
  difficulty: string;

  @ManyToMany(() => TrainingSession, (session) => session.techniques)
  trainingSession: TrainingSession[];
}