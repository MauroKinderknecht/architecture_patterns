import { Controller, Delete, Get, Post, Put } from "@nestjs/common";
import { Item } from "src/repository/item.repository";
import { ItemService } from "src/service/item.service";

@Controller('items')
export class ItemController {
    constructor(private readonly service: ItemService) {}

    @Get()
    async getAllItems() {
        return await this.service.getAllItems();
    }

    @Get(':id')
    async getItemById(id: number) {
        return await this.service.getItemById(id);
    }

    @Post()
    async createItem(item: Omit<Item, 'id'>) {
        return await this.service.createItem(item);
    }

    @Put(':id')
    async updateItem(id: number, item: Partial<Item>) {
        return await this.service.updateItem(id, item);
    }

    @Delete(':id')
    async deleteItem(id: number) {
        return await this.service.deleteItem(id);
    }
}