import React, { FC,useState,useEffect,Component } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import {
  Header,
  Page,
  pageTheme,
  Content,
 } from '@backstage/core';
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntUser } from '../../api';
import { DefaultApi } from '../../api/apis';


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      Group 18 SA-63{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}


const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
    //display: 'flex',
   // flexDirection: 'column',
    //alignItems: 'center',
    //align: 'center',
    fontSize: '18px',
  },
  avatar: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(84),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(2, 0, 2),
  },
  textField: {
    width: 350 ,
    marginLeft:7,
    marginRight:-7,
   },

}));
class Login extends Component{
   constructor(props){
    super(props);
    this.state = {
      email: '',
      password: ''
    };
}
}

const SignIn: FC<{}> = () => {

  const [users, setUsers] = useState<EntUser[]>([]);
  const http = new DefaultApi();
  function login(){

  }

  const classes = useStyles();
  return (
    <Page theme={pageTheme.tool}>

<Header
       title="Login" type="Repairing computer systems"> 
     </Header>

     <Content align="center">
  <div className={classes.paper}> <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar></div>
     <div className={classes.paper}><strong>พนักงานศูนย์บริการแจ้งซ่อมคอมพิวเตอร์</strong></div>
     <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
               // value={personalid}
                //onChange={handlePersonalIDChange}
              />
 <div></div>
      <TextField className={classes.textField}
    //          style={{ width: 500 ,marginLeft:7,marginRight:-7}}
                variant="outlined"
                margin="normal"
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
               // value={personalid}
                //onChange={handlePersonalIDChange}  
              />
<div> <Button
            type="submit"
            href="/home"
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Sign In
          </Button></div>
     </Content>
    </Page>
  );
};

export default SignIn;
