package ipaddresses_validate

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Version 1.0.3
//version 1.0.1 also updated

type IPAddressesWithLoginAndPassword struct {
	IPAddress string
	Login     string
	Password  string
}

func CheckSingleIp(IPaddr string) (err error) {
	regExString := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
	Re := regexp.MustCompile(regExString)
	iind := Re.FindStringIndex(IPaddr)
	if iind != nil && len(iind) > 0 {
		DigitVal := strings.Split(IPaddr, ".")
		for DigitIndex, DigitOne := range DigitVal {
			convval, converr := strconv.Atoi(DigitOne)
			if converr != nil {
				return errors.New("Invalid IP address")
			} else {
				if DigitIndex == 0 {
					if convval == 0 {
						return errors.New("Invalid IP address")
					}
				}
				if convval > 255 {
					return errors.New("Invalid IP address")
				}
			}
		}

		return nil
	} else {
		return errors.New("Invalid IP address")
	}
}

func MakeListIPAddresses(IPaddrs string) (IPaddrsList []string, err error) {
	var ipsplitted []string
	IPaddrs = strings.TrimSpace(IPaddrs)
	ipsplitted2 := strings.Split(IPaddrs, ",")
	for _, vars := range ipsplitted2 {
		regExString := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
		Re := regexp.MustCompile(regExString)
		iind1 := Re.FindStringIndex(vars)

		regExString = "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}-[0-9]{1,3}$"
		Re = regexp.MustCompile(regExString)
		iind2 := Re.FindStringIndex(vars)

		if iind1 == nil && iind2 == nil {
			return ipsplitted, errors.New("Invalid IP address")
		}

		if strings.Index(vars, "-") != -1 {
			//Диапазон указывается например так: 192.168.0.10-220
			vst := strings.Split(vars, "-")

			EndRange, EndRangeErr := strconv.Atoi(vst[1])
			vsi := strings.Split(vst[0], ".")
			firstipinrange_lastoctet, firstipinrange_lastoctetErr := strconv.Atoi(vsi[3])
			if firstipinrange_lastoctetErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			}

			if EndRangeErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			} else {
				if EndRange == 0 || EndRange > 255 || EndRange <= firstipinrange_lastoctet {
					return ipsplitted, errors.New("Invalid IP address")
				}
			}

			IPPrefix := vsi[0] + "." + vsi[1] + "." + vsi[2] + "."
			FirstLoctet, _ := strconv.Atoi(vsi[len(vsi)-1])
			LastLoctet, _ := strconv.Atoi(vst[1])
			for vsit := FirstLoctet; vsit <= LastLoctet; vsit++ {
				ip := IPPrefix + strconv.Itoa(vsit)
				ipsplitted = append(ipsplitted, ip)
			}
		} else {
			ipsplitted = append(ipsplitted, vars)
		}
	}
	return ipsplitted, nil
}

func MakeListIPAddressesWithLoginPassword(IPaddrs string) (IPaddrsList []IPAddressesWithLoginAndPassword, err error) {
	var ipsplitted []IPAddressesWithLoginAndPassword
	IPaddrs = strings.TrimSpace(IPaddrs)
	ipsplitted2 := strings.Split(IPaddrs, ",")
	for _, varPip := range ipsplitted2 {
		LoginPart := ""
		PasswordPart := ""
		AddressPart := varPip
		if strings.Contains(varPip, "@") {
			NestedSplittedSt := strings.Split(varPip, "@")
			if len(NestedSplittedSt) == 2 {
				if strings.Contains(NestedSplittedSt[0], ":") {
					NestedSplittedSt2 := strings.Split(NestedSplittedSt[0], ":")
					if len(NestedSplittedSt2) == 2 {
						LoginPart = NestedSplittedSt2[0]
						PasswordPart = NestedSplittedSt2[1]
						AddressPart = NestedSplittedSt[1]
					} else {
						return ipsplitted, errors.New("Invalid IP user password ipaddress string")
					}
				} else {
					return ipsplitted, errors.New("Invalid IP user password ipaddress string")
				}
			}
		}
		regExString := "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}$"
		Re := regexp.MustCompile(regExString)
		iind1 := Re.FindStringIndex(AddressPart)

		regExString = "^(?:[0-9]{1,3}\\.){3}[0-9]{1,3}-[0-9]{1,3}$"
		Re = regexp.MustCompile(regExString)
		iind2 := Re.FindStringIndex(AddressPart)

		if iind1 == nil && iind2 == nil {
			return ipsplitted, fmt.Errorf("Invalid IP address: %s", AddressPart)
		}

		if strings.Index(AddressPart, "-") != -1 {
			//Диапазон указывается например так: 192.168.0.10-220
			vst := strings.Split(AddressPart, "-")

			EndRange, EndRangeErr := strconv.Atoi(vst[1])
			vsi := strings.Split(vst[0], ".")
			firstipinrange_lastoctet, firstipinrange_lastoctetErr := strconv.Atoi(vsi[3])
			if firstipinrange_lastoctetErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			}

			if EndRangeErr != nil {
				return ipsplitted, errors.New("Invalid IP address")
			} else {
				if EndRange == 0 || EndRange > 255 || EndRange <= firstipinrange_lastoctet {
					return ipsplitted, errors.New("Invalid IP address")
				}
			}

			IPPrefix := vsi[0] + "." + vsi[1] + "." + vsi[2] + "."
			FirstLoctet, _ := strconv.Atoi(vsi[len(vsi)-1])
			LastLoctet, _ := strconv.Atoi(vst[1])
			for vsit := FirstLoctet; vsit <= LastLoctet; vsit++ {
				ip := IPPrefix + strconv.Itoa(vsit)
				ipsplitted = append(ipsplitted, IPAddressesWithLoginAndPassword{IPAddress: ip, Login: LoginPart, Password: PasswordPart})
			}
		} else {
			ipsplitted = append(ipsplitted, IPAddressesWithLoginAndPassword{IPAddress: AddressPart, Login: LoginPart, Password: PasswordPart})
		}
	}
	return ipsplitted, nil
}
