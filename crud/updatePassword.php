<?php include 'db.php';

	if (isset($_POST['submit'])) {
		$username = $_POST['username'];
		$password = $_POST['password'];
		$confirm = $_POST['password2'];

		if ($confirm != $password) {
			echo "<script>alert('Passwords do not match :(');</script>";
			die();
		}


		$query = "UPDATE user SET password = '$password' WHERE username = '$username'";

		$result = mysqli_query($connection, $query);

		if ($result) {
			echo "<script>alert('password for $username updated!');</script>";
		}
	}

?>

<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css">
	<title>Update data stored in Database</title>
</head>
<body>
	<button class="btn btn-primary" style="float: right; margin-right: 50px;" onclick="window.location = 'index.html'">Home</button>
	<div class="container" style="margin-top: 50px;">
	    
	    <div class="col-sm-3">
	        <form action="u_for_update.php" method="post">
	            <div class="form-group">
	            	<label for="username">Username</label>
	            	<input type="text" name="username" class="form-control">
	            </div>
	            
	            <div class="form-group">
	                <label for="password">Password</label>
	            	<input type="password" name="password" class="form-control">
	            </div>	            

	            <div class="form-group">
	                <label for="password2">Confirm password</label>
	            	<input type="password" name="password2" class="form-control">
	            </div>
	            
	            <input class="btn btn-primary" type="submit" name="submit" value="Update">
	            
	        </form>
	    </div>


	
</body>
</html>