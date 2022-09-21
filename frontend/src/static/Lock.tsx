import * as React from 'react'

function SvgLock(props: React.SVGProps<SVGSVGElement>) {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 64 64"
			width="1em"
			height="1em"
			{...props}
		>
			<path d="M48.1 19v-2.4C48.1 8 41.7.8 33.4 0 28.8-.2 24.2 1.1 21 4.3c-3.2 2.9-5.1 7.2-5.1 11.5V19C9.8 20.4 5.2 25.7 5.2 32.2v19C5.2 58.4 11.1 64 18 64h27.8c7.2 0 12.8-5.9 12.8-12.8v-19c.2-6.5-4.4-11.8-10.5-13.2zM24.8 8.3c2.1-2.1 5.1-2.9 8.3-2.7 5.4.5 9.6 5.4 9.6 11.2v1.9H21.3v-2.9c0-2.9 1.3-5.6 3.5-7.5zm28.6 43.1c0 4.3-3.5 7.5-7.5 7.5H18.1c-4 0-7.5-3.5-7.5-7.5v-19c0-4.5 3.7-8.3 8.3-8.3h26.2c4.5 0 8.3 3.5 8.3 8v19.3z" />
			<path d="M32 34.8c-1.6 0-2.7 1.1-2.7 2.7v10.8c0 1.6 1.1 2.7 2.7 2.7s2.7-1.1 2.7-2.7V37.5c0-1.6-1.1-2.7-2.7-2.7z" />
		</svg>
	)
}

export default SvgLock