import { createSlice } from "@reduxjs/toolkit";
import { HYDRATE } from "next-redux-wrapper";
import { AnyAction, combineReducers } from "redux";


export interface AuthState {
    email: string;
    firstName: string;
    lastName: string;
    authState: boolean;
}

const initialState: AuthState = {
    email: '',
    firstName: '',
    lastName: '',
    authState: false,
}

export interface State {
    tick: string;
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setCredentials: (state, action) => {
            state.authState = true,
            state.email = action.payload.email,
            state.firstName = action.payload.firstName,
            state.lastName = action.payload.lastName
        },
        logout: (state) => {
            state.authState = false;
            state.email = '';
            state.firstName = '';
            state.lastName = '';
        }
    },
});

export const { setCredentials, logout } = authSlice.actions;

export default authSlice.reducer;

export const selectCurrentEmail = (state : any) => state.auth.email as string;
export const selectCurrentFirstName = (state : any) => state.auth.firstName as string;
export const selectCurrentLastName = (state : any) => state.auth.lastName as string;
export const selectAuthState = (state : any) => state.auth.authState as boolean;
