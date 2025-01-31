import { Technique } from "src/modules/technique/domain/technique.entity";
import { User } from "src/modules/user/domain/user.entity";
import { Column, Entity, JoinTable, ManyToMany, ManyToOne, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class TrainingSession {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @ManyToOne(() => User, (user) => user.trainingSessions, { onDelete: 'CASCADE' })
  user: User;

  @Column()
  date: Date;

  @Column()
  duration: number; // in minutes

  @Column()
  notes: string[];

  @ManyToMany(() => Technique)
  @JoinTable()
  techniques: Technique[];
}