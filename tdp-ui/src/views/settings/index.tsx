import React from 'react';
import { ViewBase } from '../../shared/components/view-base/view-base';

const Settings = () => {
    return (
        <ViewBase
            title='Настройки'
            footerOptions={{
                buttonText: 'Сохранить',
                buttonHandler: () => alert('сохранил настройки тебе за щеку'),
                showButton: true
            }}
        >
            настроечки
        </ViewBase>
    )
}

export default Settings;
