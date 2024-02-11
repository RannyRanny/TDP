import React, {ReactNode, useState} from 'react';
import { Link } from 'react-router-dom';

interface ViewBaseProps {
	title: string;
	showGoal?: boolean;
	footerOptions: {
		showButton?: boolean;
		buttonText: string;
		buttonHandler: () => void;
	};
	children: ReactNode;
};

/**
 * Компонент-оболочка для экранов приложения
 * @param title
 * @param showGoal
 * @param footerOptions
 * @param children
 * @constructor
 */
export const ViewBase = ({
	title,
	showGoal,
	footerOptions,
	children
}: ViewBaseProps) => {
	const [goal, setGoal] = useState<string>('');

	const changeGoalHandler: HTMLInputElement['onchange'] = (e) => {
		const { value } = e.target;

		setGoal(value);
	}

	const footerButtonHandler: HTMLButtonElement['onclick'] = (e) => {
		e.preventDefault();

		if (footerOptions.buttonHandler) {
			footerOptions.buttonHandler();
		}
	}

	return (
		<section className='view-base'>
			<header className='view-base__header'>
				<div className='head-content'>
					<Link to='/'>📆</Link>
					<h1>{ title }</h1>
					<Link to='/settings'>⚙️</Link>
				</div>
				{ showGoal && (
					<div className='goal-container'>
						<input
							type='text'
							placeholder='Цель дня'
							value={ goal }
							onChange={ changeGoalHandler }
						/>
					</div>
				) }
			</header>
			<main className='view-base__main'>
				{ children }
			</main>
			<footer className='view-base__footer'>
				{ footerOptions.showButton && (
					<button
						type='button'
						onClick={ footerButtonHandler }
						className='btn btn-primary'
					>
						{ footerOptions.buttonText }
					</button>
				) }
			</footer>
		</section>
	)
}
