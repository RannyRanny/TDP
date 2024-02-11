export type TaskListItemProps = {
    /**
     * Идентификатор задачи
     */
    taskId: string;

    /**
     * Название задачи
     */
    title: string;

    /**
     * Флаг выполнения
     */
    done: boolean;

    /**
     * Час в который должна быть выполнена задача
     */
    taskHour: string;

    /**
     * Мета инфа
     */
    meta: {
        epicId: string;
    }

    /**
     * Хэндлер на чекбокс
     */
    onChecked?: (taskId: string) => void;

    /**
     * Идентификатор для тест-систем
     */
    dataTestId: string;
}
