import { BeltProgress } from "src/modules/belt/domain/belt-progress.entity";
import { TrainingSession } from "src/modules/training/domain/training-session.entity";
import { Column, Entity, OneToMany, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class User {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  name: string;

  @Column({ unique: true })
  email: string;

  @Column()
  password: string;

  @OneToMany(() => BeltProgress, (beltProgress) => beltProgress.user)
  beltProgress: BeltProgress[];

  @OneToMany(() => TrainingSession, (trainingSession) => trainingSession.user)
  trainingSessions: TrainingSession[];

  @Column({ type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
  createdAt: Date;
}