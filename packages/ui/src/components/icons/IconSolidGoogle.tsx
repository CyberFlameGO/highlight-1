import React, { useId } from 'react'
import { IconProps } from './types'

export const IconSolidGoogle = ({ size = '1em', ...props }: IconProps) => {
	const colorId1 = useId()

	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width={size}
			height={size}
			fill="none"
			viewBox="0 0 20 20"
			focusable="false"
			{...props}
		>
			<g clipPath={`url(#${colorId1})`}>
				<path
					fill="#4285F4"
					d="M18.215 10.368a7.05 7.05 0 0 0-.172-1.669h-7.665v3.03h4.5c-.092.753-.581 1.886-1.67 2.648l-.015.101 2.423 1.878.168.017c1.542-1.424 2.431-3.52 2.431-6.005Z"
				/>
				<path
					fill="#34A853"
					d="M10.378 18.35c2.204 0 4.055-.726 5.406-1.977l-2.576-1.996c-.69.48-1.614.816-2.83.816-2.158 0-3.99-1.424-4.644-3.392l-.096.008-2.52 1.95-.032.092a8.157 8.157 0 0 0 7.292 4.499Z"
				/>
				<path
					fill="#FBBC05"
					d="M5.734 11.801a5.025 5.025 0 0 1-.272-1.614c0-.563.1-1.107.263-1.615l-.004-.108-2.552-1.981-.083.04a8.17 8.17 0 0 0-.871 3.664 8.17 8.17 0 0 0 .87 3.664l2.65-2.05Z"
				/>
				<path
					fill="#EB4335"
					d="M10.378 5.18c1.533 0 2.567.662 3.157 1.215l2.303-2.25c-1.414-1.314-3.256-2.122-5.46-2.122a8.157 8.157 0 0 0-7.292 4.5l2.64 2.05c.661-1.97 2.494-3.393 4.652-3.393Z"
				/>
			</g>
			<defs>
				<clipPath id={colorId1}>
					<path fill="currentColor" d="M2.215 2h16v16.373h-16z" />
				</clipPath>
			</defs>
		</svg>
	)
}
