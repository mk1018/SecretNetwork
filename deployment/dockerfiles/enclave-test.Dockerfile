FROM baiduxlab/sgx-rust:1804-1.1.2

ENV PATH="/root/.cargo/bin:$PATH"
ARG SGX_MODE=SW
ENV SGX_MODE=${SGX_MODE}
ARG FEATURES="test"
ENV FEATURES=${FEATURES}
ENV PKG_CONFIG_PATH=""
ENV LD_LIBRARY_PATH=""
#ENV MITIGATION_CVE_2020_0551=LOAD

# Set working directory for the build
WORKDIR /enclave-test/

# Add source files
COPY third_party/build third_party/build
COPY cosmwasm/ cosmwasm/
COPY Makefile Makefile
COPY api_key.txt /enclave-test/cosmwasm/packages/wasmi-runtime/
COPY spid.txt /enclave-test/cosmwasm/packages/wasmi-runtime/

RUN make vendor

COPY deployment/ci/enclave-test.sh .
RUN chmod +x enclave-test.sh

ENTRYPOINT ["/bin/bash", "enclave-test.sh"]
