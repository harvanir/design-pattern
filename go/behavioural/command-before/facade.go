package command_before

type serviceToggleFacade struct {
	client HttpClient
}

func (st *serviceToggleFacade) On() {
	// turn on android
	st.OnIbft()
	st.OnTopup()
	st.OnKyc()
}

func (st *serviceToggleFacade) Off() {
	// turn off android
	st.OffIbft()
	st.OffTopup()
	st.OffKyc()
}

func (st *serviceToggleFacade) OnIbft() {
	requestBody := "{\"mode\": \"on\", \"service\": \"ibft\"}"
	st.client.Call("http://firebase.com", requestBody)
}
func (st *serviceToggleFacade) OnTopup() {
	requestBody := "{\"mode\": \"on\", \"service\": \"topup\"}"
	st.client.Call("http://firebase.com", requestBody)
}
func (st *serviceToggleFacade) OnKyc() {
	requestBody := "{\"mode\": \"on\", \"service\": \"kyc\"}"
	st.client.Call("http://firebase.com", requestBody)
}
func (st *serviceToggleFacade) OffIbft() {
	requestBody := "{\"mode\": \"off\", \"service\": \"ibft\"}"
	st.client.Call("http://firebase.com", requestBody)
}
func (st *serviceToggleFacade) OffTopup() {
	requestBody := "{\"mode\": \"off\", \"service\": \"topup\"}"
	st.client.Call("http://firebase.com", requestBody)
}
func (st *serviceToggleFacade) OffKyc() {
	requestBody := "{\"mode\": \"off\", \"service\": \"kyc\"}"
	st.client.Call("http://firebase.com", requestBody)
}
