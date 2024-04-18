// SPDX-License-Identifier: MIT

pragma solidity 0.8.24;

/**
 * @title Library Struts
 * @dev This lib exposes Structs used within the VerificationV0 contract.
 */
library Structs {
  /**
   * @dev Service consolidates data for a verification service such as Plaid.
   *
   * @param name the name of the verification service, e.g. "Plaid"
   * @param serviceType the type of the verification service, e.g. "KYC"
   * @param coutry the ISO 3166-2 code of the juridiction for which the verification applies, e.g. "US"
   */
  struct Service {
    string name;
    string serviceType;
    string country;
  }

  /**
   * @dev Service is DTO consolidating data for am individual verification binding a VerificationV0 tokenId to a verification service.
   *
   * @param tokenId the id of the token
   * @param service the name of the verification service
   */
  struct Verification {
    uint256 tokenId;
    string service;
  }

  /**
   * @dev VerificationV0 is a DTO consolidating data for am individual verification.
   * See VerificationV0::verifications(address owner)
   *
   * @param tokenId the id of the token
   * @param service the name of the service, e.g. "Plaid"
   * @param timestamp the UNIX timestamp of when the verification was issued
   */
  struct VerificationV0 {
    uint256 tokenId;
    string service;
    uint256 timestamp;
  }
}
