#!/bin/bash
./tpm-iso20022-cli gen \
     --xsdfiles "~/iso-20022/schemas/pain.001.001.03.xsd" \
     -o "~/iso-20022/messages/pain/001.001.03" \
     --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03" \
     --xsdtPkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"

./tpm-iso20022-cli gen \
     --xsdfiles "~/iso-20022/schemas/pain.002.001.03.xsd" \
     -o "~/iso-20022/messages/pain/002.001.03" \
     --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03" \
     --xsdtPkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"

./tpm-iso20022-cli gen \
     --xsdfiles "~/iso-20022/schemas/pain.008.001.02.xsd" \
     -o "~/iso-20022/messages/pain/008.001.02" \
     --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02" \
     --xsdtPkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"

./tpm-iso20022-cli gen \
     --xsdfiles "~/iso-20022/schemas/pain.013.001.07.xsd" \
     -o "~/iso-20022/messages/pain/013.001.07" \
     --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/013.001.07" \
     --xsdtPkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"

./tpm-iso20022-cli gen \
     --xsdfiles "~/iso-20022/schemas/camt.053.001.02.xsd" \
     -o "~/iso-20022/messages/camt/053.001.02" \
     --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/camt/053.001.02" \
     --xsdtPkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/xsdt"