import snappi
import pandas

api = snappi.Api()
config_state = api.control_state.config_state
config_state.state = snappi.CONFIGSTATE_SET
config = config_state.config
config.ports.append(name='Tx Port', location='10.36.74.26;02;13')
config.ports.append(name='Rx Port', location='10.36.74.26;02;14')
flow = config.flows.append(name='Tx -> Rx Flow')
flow.tx_rx.port_tx_rx.set(tx_port_name=config.ports[0].name,
                          rx_port_name=config.ports[1].name)
flow.packet.append(choice=snappi.FLOWHEADER_ETHERNET) \
    .append(choice=snappi.FLOWHEADER_VLAN) \
    .append(choice=snappi.FLOWHEADER_IPV4) \
    .append(choice=snappi.FLOWHEADER_TCP)
flow.size.fixed = 128
flow.duration.fixed_packets = 10000
flow.rate.pps = 1000
api.set_state(api.control_state)

control_state.flow_transmit_state.state = snappi.FLOWTRANSMITSTATE_START
api.set_state(control_state)

request = api.result_portrequest.column_names = [
    snappi.PORTREQUEST_NAME, snappi.PORTREQUST_LOCATION,
    snappi.PORTREQUEST_FRAMES_TX, snappi.PORTREQUEST_FRAMES_RX
])
while True:
    results = api.get_port_results(request)
    df = pandas.DataFrame.from_dict(results)
    print(df)
    if df.frames_tx.sum() >= config.flows[0].duration.fixed_packets:
        break