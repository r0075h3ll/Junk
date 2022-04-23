<?php include "db.php";

	//this code is vulnerable to SQL Injection

	if (isset($_POST['submit'])) {
		$user = $_POST['username'];
		$pass = $_POST['password'];


		$connection = mysqli_connect('localhost', 'root', '', 'loginapp'); //making connection to the database

		if ($connection) {
		} else {
			die();
		}

		$check = "SELECT * FROM user WHERE username = '$user'";
		$checkUsername = mysqli_query($connection, $check);

		while ($row = mysqli_fetch_assoc($checkUsername)) {
			if ($user == $row['username']) {
				echo "<script>alert('Username not available');</script>";
				die();
			}
		}

		$query = "INSERT INTO user(username, password) "; 
		$query .= "VALUES ('$user', '$pass')";

		$result = mysqli_query($connection, $query);

		if (!$result) {
			die();
		}
	}
?>


<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet">
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"></script>
</head>
<body>
	<div class="m-4">
    <!-- Success Alert -->
    <div class="alert alert-success alert-dismissible fade show">
        Congrats! Account created successfuly!
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
    </div>
	
</body>
</html>