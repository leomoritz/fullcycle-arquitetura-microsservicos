/**
 * Database DB configuration module
 */
import { Injectable } from '@nestjs/common';
import { TypeOrmModuleOptions } from '@nestjs/typeorm';

@Injectable()
export class DatabaseConfig {
  public get DBConfig(): TypeOrmModuleOptions {
    return {
      type: 'mysql',
      host: process.env.DB_HOST || 'localhost',
      port: parseInt(process.env.DB_PORT || '3306', 10) || 3306,
      username: process.env.DB_USER || 'root',
      password: process.env.DB_PASSWORD || 'root',
      database: process.env.DB_NAME || 'balance_db',
      entities: [__dirname + '/../../**/*.entity{.ts,.js}'],
      synchronize: true,
    };
  }
}