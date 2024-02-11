import React from 'react';
import { TaskListItemProps } from '../typings';

export const TaskListItemComponent = ({
    taskId,
    title,
    done,
    taskHour,
    onChecked,
    ...props
}: TaskListItemProps) => {
    const taskIdLabel = `taskId-${ taskId }`;
    const changeHandler = () => {
        if (onChecked) {
            onChecked(taskId);
        }
    }

    return (
        <div key={ taskId } className='task-list-item' data-test-id={ props.dataTestId }>
            <span>{ taskHour }</span>
            <span>{ title }</span>
            <div>
                <label htmlFor={ taskIdLabel } />
                <input
                    type='checkbox'
                    checked={ done }
                    onChange={ changeHandler }
                    id={ taskIdLabel }
                />
            </div>
        </div>
    )
}
