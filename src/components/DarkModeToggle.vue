<template>
	<div>
		<svg
			id="eg5pwu10v571"
			xmlns="http://www.w3.org/2000/svg"
			xmlns:xlink="http://www.w3.org/1999/xlink"
			viewBox="0 0 796 796"
			shape-rendering="geometricPrecision"
			text-rendering="geometricPrecision"
		>
			<circle
				id="eg5pwu10v573"
				r="268"
				transform="matrix(1 0 0 1 398.18936602511235 397.81063397507751)"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<g id="eg5pwu10v574_to" transform="translate(260,0)">
				<circle
					id="eg5pwu10v574"
					r="268"
					transform="translate(268,268)"
					fill="rgb(255,255,255)"
					stroke="none"
					stroke-width="1"
				/>
			</g>
			<rect
				id="eg5pwu10v575"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 258.93798828125000 627.58801269531250)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v576"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 462.01098632812500 658.67498779296875)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v577"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 137.32499694824219 462.01098632812500)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v578"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 168.41200256347656 258.93701171875000)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v579"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.88300001621246 0.46900001168251 -0.46900001168251 0.88300001621246 557.99102783203125 65.15000152587891)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v5710"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.95630475596304 -0.29237170472274 0.29237170472274 0.95630475596304 275.76998901367188 49.50899887084961)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v5711"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(0.29237170472274 0.95630475596304 -0.95630475596304 0.29237170472274 746.49102783203125 275.76998901367188)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<rect
				id="eg5pwu10v5712"
				width="30"
				height="101"
				rx="15"
				ry="15"
				transform="matrix(-0.46900001168251 0.88300001621246 -0.88300001621246 -0.46900001168251 730.84997558593750 557.99102783203125)"
				opacity="0"
				fill="rgb(139,139,139)"
				stroke="none"
				stroke-width="1"
			/>
			<clipPath id="eg5pwu10v5713">
				<rect
					id="eg5pwu10v5714"
					width="796"
					height="796"
					rx="0"
					ry="0"
					fill="rgb(0,0,0)"
					stroke="none"
					stroke-width="1"
				/>
			</clipPath>
		</svg>

		<input type="checkbox" id="toggle" class="checkbox" />
		<label for="toggle" class="switch"></label>
	</div>
</template>
<script>
export default {
	name: "DarkModeToggle",
};
</script>
<style scoped>

.switch {
	position: relative;
	display: inline-block;
	width: 40px;
	height: 20px;
	background-color: #eee;
	border-radius: 20px;
}

.switch::after {
	content: "";
	position: absolute;
	width: 18px;
	height: 18px;
	border-radius: 50%;
	background-color: white;
	top: 1px;
	left: 1px;
	transition: all 0.3s;
}

.checkbox:checked + .switch::after {
	left: 20px;
}

.checkbox:checked + .switch {
	background-color: #7983ff;
}

.checkbox {
	display: none;
}

#eg5pwu10v573 {
	animation: eg5pwu10v573__rd 1000ms linear 1 normal forwards;
}

@keyframes eg5pwu10v573__rd {
	0% {
		r: 268px;
	}
	100% {
		r: 202px;
	}
}
#eg5pwu10v574_to {
	animation: eg5pwu10v574_to__to 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v574_to__to {
	0% {
		transform: translate(260px, 0px);
	}
	100% {
		transform: translate(560.796528px, -311.186989px);
	}
}
#eg5pwu10v574 {
	animation: eg5pwu10v574_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v574_c_o {
	0% {
		opacity: 1;
	}
	50% {
		opacity: 0;
	}
	100% {
		opacity: 0;
	}
}
#eg5pwu10v575 {
	animation: eg5pwu10v575_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v575_c_o {
	0% {
		opacity: 0;
	}
	30% {
		opacity: 1;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v576 {
	animation: eg5pwu10v576_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v576_c_o {
	0% {
		opacity: 0;
	}
	30% {
		opacity: 1;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v577 {
	animation: eg5pwu10v577_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v577_c_o {
	0% {
		opacity: 0;
	}
	30% {
		opacity: 1;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v578 {
	animation: eg5pwu10v578_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v578_c_o {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v579 {
	animation: eg5pwu10v579_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v579_c_o {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v5710 {
	animation: eg5pwu10v5710_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v5710_c_o {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v5711 {
	animation: eg5pwu10v5711_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v5711_c_o {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}
#eg5pwu10v5712 {
	animation: eg5pwu10v5712_c_o 1000ms linear 1 normal forwards;
}
@keyframes eg5pwu10v5712_c_o {
	0% {
		opacity: 0;
	}
	100% {
		opacity: 1;
	}
}
</style>
