<?php include 'db.php';

	if (isset($_POST['submit'])) {
		$username = $_POST['username'];
		$query = "DELETE FROM user WHERE username = '$username'";
		$check = "SELECT username FROM user";

		$checkUser = mysqli_query($connection, $check);

		while ($row = mysqli_fetch_assoc($checkUser)) {
			if (in_array($username, $row)) {
				$flag = 1;
				break;
			}
		}

		if ($flag) {
			$result = mysqli_query($connection, $query);
			if (!$result) {
				die();
			} else {
				echo "<script>alert('Account deleted')</script>";
			}
		} else {
			echo "<script>alert('Account doesnot exist!')</script>";
		}

	}

?>

<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css">
</head>

<body>
	<div class="container" style="margin-top: 50px;">
	    <div class="col-sm-3">
	        <form action="deleteAccount.php" method="post">
	            <div class="form-group">
	            	<label for="username">Username</label>
	            	<input type="text" name="username" class="form-control">
	            </div>
	            
	            <input class="btn btn-primary" type="submit" name="submit" value="Delete">
	            
	        </form>
	    </div>
	</div>
</body>
</html>