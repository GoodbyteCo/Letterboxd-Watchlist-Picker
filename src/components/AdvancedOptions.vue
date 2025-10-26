---
interface Props {
	legend: string
	/** the query param name, used in URLSearchParams */
	name: string
	/** the available options in {"label": "value"} form */
	options: {[label: string]: string}
	/** currently selected options */
	selected: string[]
}

const props = Astro.props
const options = Object.entries(props.options)
---

<div>
	<fieldset>
		<legend>
			<b>{props.legend}</b>
		</legend>
		<ul>
			{options.map(([label, value]) => (
				<li>
					<label>
						<input
							type="checkbox"
							name={props.name}
							value={value}
							checked={props.selected.includes(value)}
						/>
						<span>
							{label}
						</span>
					</label>
				</li>
			))}
		</ul>
	</fieldset>
</div>

<style>
	ul, li
	{
		display: inline;
		list-style: none;
		padding: 0;
	}

	label
	{
		display: inline-flex;
		flex-direction: row;
		align-items: center;
		gap: 0.5ch;

		background-color: var(--off-white);
		color: var(--black);
		margin: 0.8ch 0.5ch;
		padding: 5px 1ch;
		border-radius: 6px;
		cursor: pointer;
		user-select: none;
	}

	.dark label
	{
		opacity: 0.6;
	}

	fieldset
	{
		border: none;
		margin-block: 10px;
	}
</style>
