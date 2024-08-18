<script>
	let csrfToken = "";

	// TODO: ページネーションを実装
	// TODO: 別の関数（方法）でCSRFトークンを取得
	const getTodos = async () => {
	    try {
	        const res = await fetch("http://127.0.0.1/api/todos?limit=20", { credentials: 'include' });

	        if (!res.ok) {
	            throw new Error(`HTTP error! status: ${res.status}`);
	        }

			csrfToken = res.headers.get('X-CSRF-Token');

	        const data = await res.json();
	        return data;
	    } catch (error) {
	        console.error("Failed to fetch todos:", error);
	        return [];
	    }
	};

	let newTodo = "";
	let todos = [];

	$: todos = getTodos();

	const addTodo = async (event) => {
		event.preventDefault();

		if (newTodo.trim()) {
			const response = await fetch("http://127.0.0.1/api/todos", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					"X-CSRF-Token": csrfToken,
				},
				credentials: 'include',
				body: JSON.stringify({ title: newTodo }),
			});

			if (response.ok) {
				newTodo = "";
				todos = await getTodos();
			} else {
				console.error("Failed to add todo:", response.status);
			}
		}
	};

	const deleteTodo = async (todoID) => {
		const response = await fetch(`http://127.0.0.1/api/todos/${todoID}/delete`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				"X-CSRF-Token": csrfToken,
			},
			credentials: 'include'
		});

		if (response.ok) {
			newTodo = "";
			todos = await getTodos();
		} else {
			console.error("Failed to add todo:", response.status);
		}
	}
</script>

<div class="main">
	<form class="form" on:submit={addTodo}>
		<input
			class="text-input"
			type="text"
			placeholder="TODOを入力"
			bind:value={newTodo}
		/>
	</form>

	{#await todos}
		<p>Loading...</p>
	{:then todos}
		<ul class="todo-items">
			{#each todos as {todo_id, title}}
				<li id={todo_id} class="todo-item">
					{title}
					<form class="delte-form" on:submit={deleteTodo(todo_id)}>
						<button class="delete-btn" type="submit"><i class="fa-solid fa-xmark"></i></button>
					</form>
				</li>
			{/each}
		</ul>
	{/await}
</div>


<style>
	input {
		padding: 0;
		margin: 0;
	}
	button {
		padding: 0;
		margin: 0;
	}
	ul {
		padding: 0;
		list-style: none;
	}
	.main {
		max-width: 600px;
		padding: 30px 20px 0px 20px;
		box-sizing: border-box;
		width: 100%;
		margin: 0 auto;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
	}
	.form {
		width: 100%;
	}
	.text-input {
		background-color: rgba(172, 189, 202, 0.717);
		border: none;
		font-size: 18px;
		padding: 10px 20px;
		box-sizing: border-box;
		border-radius: 7px;
		outline: none;
		width: 100%;
		letter-spacing: 1px;
	}
	.todo-items {
		width: 100%;
	}
	.todo-item {
		font-weight: 600;
		display: flex;
		align-items: center;
		justify-content: space-between;
		height: 40px;
		font-size: 20px;
		margin-top: 10px;
	}
	.todo-item:nth-child(1) {
		margin-top: 0;
	}
	.delete-btn {
		outline: none;
		border: none;
		background-color: white;
		color: rgb(45, 69, 165);
		cursor: pointer;
	}
</style>

